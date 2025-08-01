<template>
  <el-container>
    <!-- 搜索栏 -->
    <el-header>
      <div class="header-row">
        <el-input
          v-model="searchKeyword"
          clearable
          placeholder="搜索用户名、凭证或备注"
          style="width: 300px; margin-right: 10px"
        />

        <el-upload
          :http-request="handleImportUsers"
          :limit="1"
          accept=".zip,.json"
        >
          <template #trigger>
            <el-button ref="innerBtn" style="display: none"></el-button>
          </template>
        </el-upload>

        <el-button-group class="ml-4">
          <el-button type="warning" plain @click="handleImportUsersClick"
            >导入用户
          </el-button>
          <el-button type="warning" plain @click="handleExportUsers()"
            >导出用户
          </el-button>
          <el-popconfirm
            title="确定清空客户端配置吗？"
            @confirm="handleDeleteAll"
          >
            <template #reference>
              <el-button type="danger" plain>清空用户</el-button>
            </template>
          </el-popconfirm>
          <el-button
            type="primary"
            plain
            @click="showDialog('add', createNewUser())"
            >新增用户
          </el-button>
          <el-popconfirm
            title="确定删除吗？"
            v-if="selectData.length !== 0"
            @confirm="handleDeleteUsers"
          >
            <template #reference>
              <el-button type="danger" plain>删除用户</el-button>
            </template>
          </el-popconfirm>
          <el-button type="success" plain @click="handleRefresh"
            >刷新
          </el-button>
          <el-popconfirm
            title="Are you sure to upload config?"
            @confirm="handleUploadCloud"
            @cancel="cloudApiForm.isShow = true"
          >
            <template #reference>
              <el-button type="info" plain>上传云端</el-button>
            </template>
          </el-popconfirm>
          <el-popconfirm
            title="Are you sure to upgrade config?"
            @confirm="handleUpgradeCloud"
            @cancel="cloudApiForm.isShow = true"
          >
            <template #reference>
              <el-button type="info" plain>同步云端</el-button>
            </template>
          </el-popconfirm>

          <!--          <el-button type="info" plain @click="drawer = true">测试</el-button>-->
        </el-button-group>
      </div>
    </el-header>

    <!-- 表格 -->
    <el-main>
      <el-table
        :data="paginatedTableData"
        style="width: 100%"
        @selection-change="handleSelectionChange"
        class="custom-border-table"
        :border="true"
        :cell-style="{ padding: mobileLayout ? '4px' : '8px' }"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column type="expand">
          <template #default="props">
            <div m="4">
              <p m="t-0 b-2">Frp连接ID: {{ props.row.id }}</p>
              <p m="t-0 b-2">Frp连接凭证: {{ props.row.token }}</p>
              <p m="t-0 b-2">允许端口: {{ props.row.ports }}</p>
              <p m="t-0 b-2">允许域名: {{ props.row.domains }}</p>
              <p m="t-0 b-2">允许子域名: {{ props.row.subdomains }}</p>
            </div>
          </template>
        </el-table-column>
        <el-table-column prop="comment" label="备注" />
        <el-table-column prop="user" label="名称" />
        <el-table-column prop="count" label="数量" width="100">
          <template #default="{ row }">
            <el-text
              size="large"
              v-if="row.count > 0"
              class="mx-1"
              type="danger"
              >{{ row.count }}
            </el-text>
          </template>
        </el-table-column>
        <el-table-column prop="enable" label="状态">
          <template #default="{ row }">
            <el-tag :type="row.enable ? 'success' : 'danger'">
              {{ row.enable ? '启动' : '禁用' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="400">
          <template #default="{ row }">
            <el-button-group class="ml-4">
              <el-button
                plain
                :type="row.enable ? 'danger' : 'success'"
                @click="showDialog('ToggleStatus', row)"
                size="small"
              >
                {{ row.enable ? '禁用' : '启用' }}
              </el-button>
              <el-button
                type="primary"
                plain
                @click="showDialog('update', row)"
                size="small"
                >编辑
              </el-button>
              <el-button
                type="primary"
                plain
                @click="handleDelete(row)"
                size="small"
                >删除
              </el-button>
              <el-button
                type="primary"
                plain
                @click="handleClientDialog(row)"
                size="small"
                >生成客户端
              </el-button>
              <el-button
                type="primary"
                plain
                size="small"
                v-if="row.count > 0"
                @click="handleGotoClientList(row)"
                >查看客户端
              </el-button>
            </el-button-group>
          </template>
        </el-table-column>
      </el-table>

      <!-- 分页 -->
      <el-pagination
        style="margin-top: 20px"
        background
        layout="prev, pager, next"
        :total="filteredTableData.length"
        :page-size="pageSize"
        :current-page="currentPage"
        :pager-count="mobileLayout ? 3 : 7"
        @current-change="handlePageChange"
      />
    </el-main>

    <!-- 新增用户弹窗 -->
    <el-dialog v-model="dialogVisible" title="新增用户" width="500px">
      <el-form
        ref="userRuleFormRef"
        :rules="userRules"
        :model="newUserForm"
        label-width="100px"
      >
        <el-form-item label="ID" prop="id">
          <el-input v-model="newUserForm.id" placeholder="请输入ID" disabled>
            <template #append v-if="!newUserForm.editable">
              <el-button @click="handleRandUser">随机</el-button>
            </template>
          </el-input>
        </el-form-item>
        <el-form-item label="凭证" prop="token">
          <el-input
            disabled
            v-model="newUserForm.token"
            placeholder="请输入Token凭证(meta_token)"
          />
        </el-form-item>
        <el-form-item label="名称" prop="user">
          <el-input
            v-model="newUserForm.user"
            placeholder="请输入名称(user)"
            :disabled="newUserForm.editable"
          >
          </el-input>
        </el-form-item>
        <el-form-item label="备注" prop="comment">
          <el-input v-model="newUserForm.comment" placeholder="请输入备注" />
        </el-form-item>
        <el-form-item label="允许端口">
          <el-input
            :rows="2"
            type="textarea"
            v-model="newUserForm.ports"
            placeholder="请输入允许使用的端口，如：8081,9000-9100"
          />
        </el-form-item>
        <el-form-item label="允许域名">
          <el-input
            :rows="2"
            type="textarea"
            v-model="newUserForm.domains"
            placeholder="请输入允许使用的域名，如：web01.domain.com,web02.domain.com"
          />
        </el-form-item>
        <el-form-item label="允许子域名">
          <el-input
            :rows="2"
            type="textarea"
            v-model="newUserForm.subdomains"
            placeholder="请输入允许使用的子域名，如：web01,web02"
          />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="handleDialogCancel">取消</el-button>
        <el-button
          type="primary"
          @click="
            submitForm(userRuleFormRef, () => {
              handleDialogConfirm()
            })
          "
          >确定
        </el-button>
      </template>
    </el-dialog>

    <!-- 生成客户端弹窗 -->
    <el-dialog
      v-model="genClientDialogVisible"
      title="生成客户端"
      width="500px"
    >
      <el-form label-width="130px">
        <el-form-item label="Frps服务地址：">
          <el-input
            v-model="clientForm.serverAddr"
            placeholder="请输入Frps服务地址"
          />
        </el-form-item>
        <el-form-item label="Frps服务端口：">
          <el-input-number
            controls-position="right"
            v-model="clientForm.serverPort"
            placeholder="请输入Frps服务端口"
          />
        </el-form-item>
        <el-form-item label="Frps Admin端口：">
          <el-input-number
            controls-position="right"
            v-model="clientForm.serverAdminPort"
            placeholder="请输入Frps Admin端口"
          />
        </el-form-item>

        <el-form-item label="操作系统/架构" v-if="options.length > 0">
          <el-cascader
            :options="options"
            clearable
            @change="handleOptionChange"
            v-model="clientForm.ops"
            placeholder="请选择"
          />
        </el-form-item>

        <el-form-item label="客户端下载地址：" v-if="options.length <= 0">
          <el-input
            v-model="clientForm.url"
            placeholder="请输入客户端下载地址"
          />
        </el-form-item>

        <el-divider content-position="left">
          <el-button
            type="text"
            @click="
              clientForm.webserver.showAdvanced =
                !clientForm.webserver.showAdvanced
            "
          >
            <span>{{
              clientForm.webserver.showAdvanced ? '收起' : 'frpc dashborad配置'
            }}</span>
          </el-button>
        </el-divider>

        <transition name="fade">
          <div v-show="clientForm.webserver.showAdvanced">
            <el-form-item label="管理地址：">
              <el-input
                v-model="clientForm.webserver.addr"
                placeholder="请输入addr"
              />
            </el-form-item>

            <el-form-item label="管理端口：">
              <el-input-number
                controls-position="right"
                v-model="clientForm.webserver.port"
                @change="clientForm.proxy.localPort = clientForm.webserver.port"
                placeholder="请输入port"
              />
            </el-form-item>
            <el-form-item label="管理用户：">
              <el-input
                v-model="clientForm.webserver.user"
                placeholder="请输入admin"
              />
            </el-form-item>
            <el-form-item label="管理密码：">
              <el-input
                v-model="clientForm.webserver.password"
                placeholder="请输入password"
              />
            </el-form-item>
          </div>
        </transition>

        <el-divider content-position="left">
          <el-button
            type="text"
            @click="clientForm.showAdvanced = !clientForm.showAdvanced"
          >
            <span>{{
              clientForm.showAdvanced ? '收起' : 'frpc 代理配置'
            }}</span>
          </el-button>
        </el-divider>

        <transition name="fade">
          <div v-show="clientForm.showAdvanced">
            <el-form-item label="代理类型：">
              <el-select
                v-model="clientForm.proxy.type"
                placeholder="代理类型选择"
                clearable
              >
                <el-option
                  v-for="item in clientForm.options"
                  :key="item.value"
                  :label="item.label"
                  :value="item.value"
                />
              </el-select>
            </el-form-item>
            <el-form-item label="代理名称：">
              <el-input
                v-model="clientForm.proxy.name"
                placeholder="请输入代理名称"
              />
            </el-form-item>
            <el-form-item label="本地地址：">
              <el-input
                v-model="clientForm.proxy.localIP"
                placeholder="请输入localIP"
              />
            </el-form-item>
            <el-form-item label="本地端口：">
              <el-input-number
                controls-position="right"
                v-model="clientForm.proxy.localPort"
                placeholder="请输入localPort"
              />
            </el-form-item>
            <el-form-item label="远程端口：">
              <el-input-number
                controls-position="right"
                v-model="clientForm.proxy.remotePort"
                placeholder="请输入remotePort"
              />
            </el-form-item>
          </div>
        </transition>
      </el-form>
      <template #footer>
        <el-upload :http-request="handleGenClientByPut" :limit="1">
          <template #trigger>
            <el-button ref="innerGenBtn" style="display: none"></el-button>
          </template>
        </el-upload>

        <el-button
          type="warning"
          plain
          @click="handleGenClick"
          v-if="!clientForm.ops"
          >上传生成
        </el-button>

        <el-button @click="downloadClientTomlFile">配置下载</el-button>
        <el-button @click="showFrpcToml">显示配置</el-button>
        <el-button
          type="danger"
          @click="downloadClientByGen"
          :loading="clientForm.loading"
          >程序生成
        </el-button>
      </template>
    </el-dialog>

    <!-- 填写云Api信息设置 -->
    <el-dialog
      v-model="cloudApiForm.isShow"
      title="云Api信息设置"
      width="500px"
    >
      <el-form label-width="130px">
        <el-form-item label="接口地址：">
          <el-input v-model="cloudApiForm.addr" placeholder="请输入api地址" />
        </el-form-item>
        <el-form-item label="授权用户：">
          <el-input v-model="cloudApiForm.user" placeholder="请输入授权用户" />
        </el-form-item>
        <el-form-item label="授权密钥：">
          <el-input v-model="cloudApiForm.pass" placeholder="请输入授权密钥" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button type="primary" @click="handleUploadCloud">确定</el-button>
      </template>
    </el-dialog>

    <el-drawer
      v-model="drawer"
      :size="isMobilePhone() ? '100%' : '50%'"
      :with-header="false"
      :direction="direction"
      :show-close="true"
    >
      <div style="margin-left: 8px">
        <el-button type="success" @click="copyText(drawertextarea)"
          >复制
        </el-button>
        <el-input
          style="margin-top: 8px"
          v-model="drawertextarea"
          autosize
          placeholder="frps configure file, can not be empty..."
          type="textarea"
        ></el-input>
      </div>
    </el-drawer>
  </el-container>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, onUnmounted, reactive, onUpdated } from 'vue'
import {
  post,
  get,
  showErrorTips,
  showWarmDialog,
  generateRandomKey,
  deepCopyJSON,
  downloadByPost,
  showLoading,
  showTips,
  DownLoadFile,
  isMobilePhone,
  copyToClipboard,
} from '../utils/utils.ts'
import { DrawerProps, ElButton, FormInstance, FormRules } from 'element-plus'
import router from '../router'
import {
  FrpcConfiguration,
  ConfigBodyData,
  TypedProxyConfig,
} from '../utils/type.ts'

const innerBtn = ref<InstanceType<typeof ElButton>>()
const innerGenBtn = ref<InstanceType<typeof ElButton>>()
// const ops = ref({})
const options = ref([])
// const isLoading = ref<boolean>(false)

const drawertextarea = ref('')
const drawer = ref(false)
const direction = ref<DrawerProps['direction']>('ltr')

// 搜索关键字
const searchKeyword = ref<string>('')
// 分页相关
const pageSize = ref<number>(10)
const currentPage = ref<number>(1)
//const filteredTableData = ref<User[]>([])

// 表格数据{user:'admin'}
// const tableData = ref<User[]>([{ user: 'admin',count: 11 }])
const tableData = ref<FrpcConfiguration[]>([])
const selectData = ref<FrpcConfiguration[]>([])
// const clientDownUrl = ref<string>()
// const serverAddr = ref<string>()
const dialogType = ref<string>()
// 新增用户弹窗相关
const dialogVisible = ref<boolean>(false)
//生成客户端弹窗相关
const genClientDialogVisible = ref<boolean>(false)
const newUserForm = ref({
  user: '',
  token: '',
  comment: '',
  ports: '',
  domains: '',
  subdomains: '',
  enable: true,
  count: 0,
  editable: false,
  id: '',
})

const clientForm = ref({
  serverAddr: '',
  serverPort: 0,
  serverAdminPort: parseInt(window.location.port, 10),
  url: '',
  ops: null,
  loading: false,
  showAdvanced: false,
  proxy: {
    type: 'tcp',
    name: '',
    localIP: '0.0.0.0',
    localPort: 6400,
    remotePort: 0,
  },
  webserver: {
    addr: '0.0.0.0',
    port: 6400,
    user: 'admin',
    password: '',
    showAdvanced: false,
  },
  options: [
    {
      label: 'tcp',
      value: 'tcp',
    },
    {
      label: 'udp',
      value: 'udp',
    },
  ],
})

const cloudApiForm = ref({
  addr: '',
  user: '',
  pass: '',
  isShow: false,
})

const userRuleFormRef = ref<FormInstance>()

const userRules = reactive<FormRules>({
  id: [
    {
      required: true,
      message: 'ID不能为空',
      trigger: 'blur',
    },
  ],
  user: [
    {
      required: true,
      message: '用户名不能为空',
      trigger: 'blur',
    },
  ],
  token: [
    {
      required: true,
      message: '凭证不能空',
      trigger: 'blur',
    },
  ],
  comment: [
    {
      required: true,
      message: '备注不能空',
      trigger: 'blur',
    },
  ],
})

const submitForm = async (
  formEl: FormInstance | undefined,
  func: () => void,
) => {
  if (!formEl) {
    console.log('formEl err')
    return
  }
  await formEl.validate((valid, fields) => {
    if (valid) {
      console.log('submit!')
      func()
    } else {
      console.log('error submit!', fields)
    }
  })
}

const copyText = (text: string) => {
  drawer.value = false
  genClientDialogVisible.value = false
  return copyToClipboard(text)
}
const showDrawer = (text: string) => {
  drawer.value = true
  drawertextarea.value = text
}
const resetForm = (formEl: FormInstance | undefined) => {
  if (!formEl) return
  formEl.resetFields()
}
const handleOptionChange = (value: any) => {
  console.log('handleOptionChange', clientForm.value.ops, value)
  //showSucessTips(JSON.stringify(node))
  if (clientForm.value.ops) {
    console.log('有值了', clientForm.value.ops)
  } else {
    console.log('没值了', clientForm.value.ops)
  }
}
// 过滤后的表格数据（根据搜索关键字）
const filteredTableData = computed<FrpcConfiguration[]>(() => {
  return tableData.value.filter(
    (data) =>
      !searchKeyword.value ||
      data.user?.includes(searchKeyword.value) ||
      data.token?.includes(searchKeyword.value) ||
      data.comment?.includes(searchKeyword.value),
  )
})

const handleGotoClientList = (row: FrpcConfiguration) => {
  //const version = inject<Ref<Version>>('version')
  //provide('FrpcConfiguration', row)
  // provide<FrpcConfiguration>('FrpcConfiguration', row)
  // frpcConfig.value = row
  // router.push({
  //   path: '/user/list',
  //   query: {
  //     profileData: JSON.stringify(row),
  //   },
  // })
  router.push({
    path: '/user/list',
    query: {
      id: row.id,
    },
  })
}

const handleDeleteUsers = () => {
  // if (selectData.value && selectData.value.length > 0) {
  //   showWarmDialog(
  //       `确定删除批量删除吗？`,
  //       () => {
  //
  //       },
  //       () => {
  //         clearVariables()
  //       },
  //   )
  // }

  const body = JSON.stringify(selectData.value)
  post('删除中...', '../api/token/del', body)
    .then((data: any) => {
      console.log(data)
      //tableData.value = tableData.value.filter((item) => item !== row)
    })
    .catch((err: any) => {
      console.log(err)
    })
    .finally(() => {
      clearVariables()
      fetchListData()
    })
}

const handleDeleteAll = () => {
  fetch('../api/token/delall', {
    method: 'POST',
  })
    .then((response) => {
      return response.json()
    })
    .finally(() => {
      clearVariables()
      fetchListData()
    })
}

const handleImportUsersClick = () => {
  console.log('配置导出中', selectData.value)
  innerBtn.value?.$el.click()
}

const handleGenClick = () => {
  console.log('handleClientGenClick')
  innerGenBtn.value?.$el.click()
}

const handleExportUsers = () => {
  console.log('配置导出中', selectData.value)

  const body = JSON.stringify(selectData.value)
  downloadByPost('配置导出中', '../api/client/user/export', body).finally(
    () => {
      genClientDialogVisible.value = false
    },
  )
}

const handleImportUsers = (options: any) => {
  const { file } = options
  const formData = new FormData()
  formData.append('file', file)
  const loading = showLoading('用户导入中...')
  // 使用 fetch 发送请求
  fetch('../api/client/user/import', {
    method: 'POST',
    body: formData,
  })
    .then((response) => {
      return response.json()
    })
    .finally(() => {
      loading.close()
      setTimeout(function () {
        window.location.reload()
      }, 1000)
    })
}
// 选择变化事件
const handleSelectionChange = (rows: FrpcConfiguration[]) => {
  selectData.value = rows
  console.log('--->', rows)
}

function createClientBodyData() {
  const bodyData: Partial<ConfigBodyData> = {
    serverAdminPort: clientForm.value.serverAdminPort,
    clientConfig: {
      serverAddr: clientForm.value.serverAddr,
      serverPort: clientForm.value.serverPort,
      proxies: [clientForm.value.proxy as TypedProxyConfig],
      webServer: {
        addr: clientForm.value.webserver.addr,
        port: clientForm.value.webserver.port,
        user: clientForm.value.webserver.user,
        password: clientForm.value.webserver.password,
      },
    },
    userConfig: {
      id: newUserForm.value.id,
      user: newUserForm.value.user,
      token: newUserForm.value.token,
      comment: newUserForm.value.comment,
      ports: toPorts(newUserForm.value.ports),
      domains: newUserForm.value.domains.split(','),
      subdomains: newUserForm.value.subdomains.split(','),
      enable: newUserForm.value.enable,
    },
  }
  console.log('--->', bodyData)
  return bodyData
}

// 调用接口创建客户端
const downloadClientByGen = () => {
  clientForm.value.loading = true
  console.log('downloadClientByGen.newUserForm', newUserForm.value)
  console.log('clientForm', clientForm.value)
  const data = createClientBodyData()
  const node = getFilePathByValue(options.value, clientForm.value.ops)
  if (node && node.filePath !== '') {
    data.binAddress = node.filePath
  } else if (clientForm.value.url !== '') {
    data.binAddress = clientForm.value.url
  } else {
    showErrorTips('生成客户端失败～')
    genClientDialogVisible.value = false
    return
  }
  downloadByPost(
    '客户端生成中',
    '../api/client/gen',
    JSON.stringify(data),
  ).finally(() => {
    genClientDialogVisible.value = false
    clientForm.value.loading = false
  })
}

const handleGenClientByPut = (options: any) => {
  const { file } = options
  const data = createClientBodyData()
  data.binAddress = clientForm.value.url
  const formData = new FormData()
  formData.append('file', file)
  formData.append('data', JSON.stringify(data))
  DownLoadFile('客户端生成中', 'PUT', '../api/client/gen', formData).finally(
    () => {},
  )
}

// 调用接口创建客户端
const downloadClientTomlFile = () => {
  const data = createClientBodyData()
  console.log('downloadClientTomlFile', clientForm.value, data)
  downloadByPost(
    '配置生成中',
    '../api/client/toml',
    JSON.stringify(data),
  ).finally(() => {
    genClientDialogVisible.value = false
  })
}

const showFrpcToml = () => {
  const data = createClientBodyData()
  fetch('../api/client/toml', {
    method: 'PUT',
    body: JSON.stringify(data),
  })
    .then((response) => {
      return response.text()
    })
    .then((text) => {
      console.log('--->', text)
      showDrawer(text)
    })
}

// const handleGenClientByPut = (options: any) => {
//   const { file } = options
//   const body = {
//     id: newUserForm.value.id,
//     user: newUserForm.value.user,
//     token: newUserForm.value.token,
//     comment: newUserForm.value.comment,
//     ports: toPorts(newUserForm.value.ports),
//     domains: newUserForm.value.domains.split(','),
//     subdomains: newUserForm.value.subdomains.split(','),
//     enable: newUserForm.value.enable,
//   }
//   const data = {
//     user: body,
//     binUrl: clientForm.value.url,
//     addr: clientForm.value.addr,
//     port: clientForm.value.port,
//     apiPort: clientForm.value.apiPort,
//     data: clientForm.value,
//     proxy: clientForm.value.proxy,
//     webserver: clientForm.value.webserver,
//   }
//   const formData = new FormData()
//   formData.append('file', file)
//   formData.append('data', JSON.stringify(data))
//   //const loading = showLoading('用户导入中...')
//   // 使用 fetch 发送请求
//   fetch('../api/client/gen', {
//     method: 'PUT',
//     body: formData,
//   })
//     .then((response) => {
//       return response.json()
//     })
//     .finally(() => {
//       loading.close()
//     })
//
//   DownLoadFile('客户端生成中', 'PUT', '../api/client/gen', formData).finally(
//     () => {},
//   )
// }
// 调用接口创建客户端
// const downloadClientByGen = () => {
//   clientForm.value.loading = true
//   console.log('downloadClientByGen.newUserForm', newUserForm.value)
//   console.log('clientForm', clientForm.value)
//   const node = getFilePathByValue(options.value, clientForm.value.ops)
//   const body = {
//     id: newUserForm.value.id,
//     user: newUserForm.value.user,
//     token: newUserForm.value.token,
//     comment: newUserForm.value.comment,
//     ports: toPorts(newUserForm.value.ports),
//     domains: newUserForm.value.domains.split(','),
//     subdomains: newUserForm.value.subdomains.split(','),
//     enable: newUserForm.value.enable,
//   }
//   if (node && node.filePath !== '') {
//     const data = {
//       binPath: node.filePath,
//       addr: clientForm.value.addr,
//       port: clientForm.value.port,
//       user: body,
//       data: clientForm.value,
//       apiPort: clientForm.value.apiPort,
//       proxy: clientForm.value.proxy,
//       webserver: clientForm.value.webserver,
//     }
//     console.log('data1', data)
//     downloadByPost(
//       '客户端生成中',
//       '../api/client/gen',
//       JSON.stringify(data),
//     ).finally(() => {
//       genClientDialogVisible.value = false
//       clientForm.value.loading = false
//     })
//   } else {
//     if (clientForm.value.url === '') {
//       showErrorTips('生成客户端失败～')
//       genClientDialogVisible.value = false
//     } else {
//       const data = {
//         binUrl: clientForm.value.url,
//         addr: clientForm.value.addr,
//         port: clientForm.value.port,
//         user: body,
//         data: clientForm.value,
//         apiPort: clientForm.value.apiPort,
//         proxy: clientForm.value.proxy,
//         webserver: clientForm.value.webserver,
//       }
//
//       console.log('data2', data)
//       downloadByPost(
//         '客户端生成中',
//         '../api/client/gen',
//         JSON.stringify(data),
//       ).finally(() => {
//         genClientDialogVisible.value = false
//         clientForm.value.loading = false
//       })
//       console.log('download----2----')
//     }
//   }
// }
//
// // 调用接口创建客户端
// const downloadClientTomlFile = () => {
//   const data = {
//     addr: clientForm.value.addr,
//     port: clientForm.value.port,
//     apiPort: clientForm.value.apiPort,
//     user: {
//       id: newUserForm.value.id,
//       user: newUserForm.value.user,
//       token: newUserForm.value.token,
//       comment: newUserForm.value.comment,
//       ports: toPorts(newUserForm.value.ports),
//       domains: newUserForm.value.domains.split(','),
//       subdomains: newUserForm.value.subdomains.split(','),
//       enable: newUserForm.value.enable,
//     },
//     data: clientForm.value,
//     proxy: clientForm.value.proxy,
//     webserver: clientForm.value.webserver,
//   }
//
//   console.log('downloadClientTomlFile', clientForm.value, data)
//   downloadByPost(
//     '配置生成中',
//     '../api/client/toml',
//     JSON.stringify(data),
//   ).finally(() => {
//     genClientDialogVisible.value = false
//   })
// }

const getFilePathByValue = (opt: any, valuePath: any) => {
  const child = opt.find((item: any) => item.value === valuePath[0])
  if (child) {
    const children = child.children
    if (children) {
      const node = children.find((item: any) => item.value === valuePath[1])
      if (node) {
        console.log(node)
        return node
      }
    }
  }
  return null
}

// 分页后的表格数据
const paginatedTableData = computed<FrpcConfiguration[]>(() => {
  const start = (currentPage.value - 1) * pageSize.value
  const end = start + pageSize.value
  return filteredTableData.value.slice(start, end)
})

// 分页切换
const handlePageChange = (page: number) => {
  currentPage.value = page
}

const handleRefresh = () => {
  fetchListData()
  fetchOptions()
}

// 配置同步云端
const handleUpgradeCloud = () => {
  console.log('handleUpgradeCloud:', cloudApiForm)
  if (cloudApiForm.value.isShow) {
    fetch('../api/config/upgrade', {
      credentials: 'include',
      method: 'post',
      body: JSON.stringify(cloudApiForm.value),
    })
      .then((res) => res.json())
      .then((json) => {
        console.log('配置备份', json)
        showTips(json.code, json.msg)
        if (json.code === 0) {
          cloudApiForm.value.isShow = false
        }
      })
      .finally(() => {
        localStorage.setItem('cloudApi', JSON.stringify(cloudApiForm.value))
        clearVariables()
        fetchListData()
      })
  } else {
    fetch('../api/config/upgrade', {
      credentials: 'include',
      method: 'get',
    })
      .then((res) => res.json())
      .then((json) => {
        console.log('配置备份', json)
        if (json.code === 100) {
          cloudApiForm.value.isShow = true
          if (json.data) {
            cloudApiForm.value.user = json.data.user
            cloudApiForm.value.pass = json.data.pass
            cloudApiForm.value.addr = json.data.addr
          }
        }
        showTips(json.code, json.msg)
      })
      .finally(() => {
        clearVariables()
        fetchListData()
      })
  }
}

// 配置上传云端
const handleUploadCloud = () => {
  console.log('handleUploadCloud:', cloudApiForm)
  if (cloudApiForm.value.isShow) {
    fetch('../api/config/upload', {
      credentials: 'include',
      method: 'post',
      body: JSON.stringify(cloudApiForm.value),
    })
      .then((res) => res.json())
      .then((json) => {
        console.log('配置备份', json)
        showTips(json.code, json.msg)
        if (json.code === 0) {
          cloudApiForm.value.isShow = false
        }
      })
      .finally(() => {
        localStorage.setItem('cloudApi', JSON.stringify(cloudApiForm.value))
        clearVariables()
        fetchListData()
      })
  } else {
    fetch('../api/config/upload', {
      credentials: 'include',
      method: 'get',
    })
      .then((res) => res.json())
      .then((json) => {
        console.log('配置备份', json)
        if (json.code === 100) {
          cloudApiForm.value.isShow = true
          if (json.data) {
            cloudApiForm.value.user = json.data.user
            cloudApiForm.value.pass = json.data.pass
            cloudApiForm.value.addr = json.data.addr
          }
        }
        showTips(json.code, json.msg)
      })
      .finally(() => {
        clearVariables()
        fetchListData()
      })
  }
}

const handleDialogCancel = () => {
  dialogVisible.value = false
  clearVariables()
  resetForm(userRuleFormRef.value)
}

// 确认新增用户
const handleDialogConfirm = () => {
  dialogVisible.value = false
  switch (dialogType.value) {
    case 'add':
      addUser()
      break
    case 'update':
      updateUser()
      break
    default:
      break
  }
}

const handleClientDialog = (row: FrpcConfiguration) => {
  genClientDialogVisible.value = true
  newUserForm.value = row
  console.log(row)
}

const showDialog = (type: string, row: FrpcConfiguration) => {
  clearVariables()
  //newUserForm.value = deepCopyJSON(row)
  //newUserForm.value = row
  if (type === 'ToggleStatus') {
    row.enable = !row.enable
    newUserForm.value = deepCopyJSON(row)
    newUserForm.value.editable = true
    updateUser()
  } else if (type === 'add') {
    newUserForm.value = deepCopyJSON(row)
    dialogVisible.value = true
    dialogType.value = type
    newUserForm.value.editable = false
  } else {
    newUserForm.value = deepCopyJSON(row)
    dialogVisible.value = true
    dialogType.value = type
    newUserForm.value.editable = true
  }
  //
}

const handleDelete = (row: FrpcConfiguration) => {
  showWarmDialog(
    `确定删除${row.user}吗？`,
    () => {
      const data = [
        {
          user: row.user,
          id: row.id,
        },
      ]
      const body = JSON.stringify(data)
      post('删除中...', '../api/token/del', body)
        .then((data: any) => {
          console.log(data)
          tableData.value = tableData.value.filter((item) => item !== row)
        })
        .catch((err: any) => {
          console.log(err)
        })
        .finally(() => {
          clearVariables()
          fetchListData()
        })
    },
    () => {
      clearVariables()
    },
  )
}

const handleRandUser = () => {
  newUserForm.value.token = `${generateRandomKey()}`
  newUserForm.value.id = `${new Date().getTime()}`
  console.log('handleRandUser', newUserForm.value)
}

const addUser = () => {
  const newData = {
    id: newUserForm.value.id,
    user: newUserForm.value.user,
    token: newUserForm.value.token,
    comment: newUserForm.value.comment,
    ports: toPorts(newUserForm.value.ports),
    domains: newUserForm.value.domains.split(','),
    subdomains: newUserForm.value.subdomains.split(','),
    enable: newUserForm.value.enable,
  }
  const body = JSON.stringify(newData)
  post('添加用户中...', '../api/token/add', body)
    .then((data: any) => {
      console.log(data)
      tableData.value.push({
        ...newUserForm.value,
        enable: true, // 默认状态为启用
      })
    })
    .catch((err: any) => {
      console.log(err)
    })
    .finally(() => {
      clearVariables()
      fetchListData()
    })
}

const updateUser = () => {
  post('更新中...', '../api/token/chg', createUser(newUserForm.value))
    .then((data: any) => {
      console.log(data)
    })
    .catch((err: any) => {
      console.log(err)
    })
    .finally(() => {
      clearVariables()
      fetchListData()
    })
}

const createUser = (row: FrpcConfiguration) => {
  const data = {
    user: row.user,
    token: row.token,
    comment: row.comment,
    count: row.count,
    ports: toPorts(row.ports),
    domains: row.domains.split(','),
    subdomains: row.subdomains.split(','),
    enable: row.enable,
    id: row.id,
  }
  return JSON.stringify(data)
}

const clearVariables = () => {
  newUserForm.value = createEmptyUser()
  dialogType.value = ''
}
const createNewUser = () => {
  return {
    user: '',
    token: `${generateRandomKey()}`,
    comment: '',
    ports: '',
    domains: '',
    subdomains: '',
    count: 0,
    enable: true,
    editable: false,
    id: `${new Date().getTime()}`,
  }
}
const createEmptyUser = () => {
  return {
    user: '',
    token: '',
    comment: '',
    ports: '',
    domains: '',
    count: 0,
    subdomains: '',
    enable: true,
    editable: false,
    id: '',
  }
}

const toPorts = (ports: string) => {
  const portArr: any[] = []
  const tempPorts = ports.split(',')
  tempPorts.forEach(function (port, index) {
    portArr[index] = port
    if (/^\d+$/.test(String(port))) {
      portArr[index] = parseInt(String(port))
    }
  })
  return portArr
}

// 响应式布局相关
const mobileLayout = ref(false)
const checkMobile = () => {
  mobileLayout.value = window.innerWidth < 768
}

// 弹窗宽度控制
const dialogWidth = ref('500px')
const updateDialogWidth = () => {
  checkMobile()
  dialogWidth.value = mobileLayout.value ? '90%' : '500px'
}

// 初始化监听
onMounted(() => {
  window.addEventListener('resize', updateDialogWidth)
  updateDialogWidth()
  clientForm.value.serverAddr = window.location.hostname
})

onUnmounted(() => {
  window.removeEventListener('resize', updateDialogWidth)
})

//watchEffect(() => {
//  filteredTableData.value = tableData.value.filter(
//      (data) =>
//          !searchKeyword.value ||
//          data.user.includes(searchKeyword.value) ||
//          data.token.includes(searchKeyword.value) ||
//          data.comment.includes(searchKeyword.value),
//  )
//})

const fetchServerData = () => {
  fetch('../api/bindinfo', { credentials: 'include' })
    .then((res) => res.json())
    .then((json) => {
      if (json.code === 0) {
        clientForm.value.serverPort = json.data.bindPort
      }
    })
    .catch(() => {
      showErrorTips('获取服务器信息失败')
    })
}

// 获取数据
const fetchListData = () => {
  get('数据请求', '../api/token/all', null).then((data: any) => {
    if (data) {
      console.log('fetchListData', data)
      const obj = JSON.parse(JSON.stringify(data))
      tableData.value = obj.map((item: any) => ({
        user: item.user,
        token: item.token,
        count: item.count,
        comment: item.comment,
        ports: item.ports.join(','),
        domains: item.domains.join(','),
        subdomains: item.subdomains.join(','),
        enable: item.enable,
        id: item.id,
      }))
    } else {
      tableData.value = []
    }
  })
}
// 获取平台数据
const fetchOptions = () => {
  get('', '../api/client/get', null).then((data: any) => {
    console.log('clients', data)
    if (data) {
      options.value = JSON.parse(JSON.stringify(data))
    } else {
      options.value = []
    }
  })
}

onUpdated(() => {
  fetchOptions()
})
onMounted(() => {
  const jsonString = localStorage.getItem('cloudApi')
  if (jsonString) {
    const obj = JSON.parse(jsonString)
    if (obj) {
      cloudApiForm.value = obj
    }
    cloudApiForm.value.isShow = false
  }
})

fetchListData()
fetchOptions()
fetchServerData()
</script>

<style scoped>
.custom-border-table {
  border: 1px solid var(--el-border-color);
  transform: translateZ(0);

  /* 斑马纹效果 */

  :deep(.el-table__row--striped) {
    background-color: var(--el-fill-color-light);
  }

  /* 单元格统一边框 */

  :deep(.el-table__cell) {
    border-right: 1px solid var(--el-border-color);
    border-bottom: 1px solid var(--el-border-color);
  }

  /* 悬浮效果 */

  :deep(.el-table__row:hover td) {
    background-color: var(--el-fill-color) !important;
  }
}

.el-header {
  display: flex;
  flex-direction: column;
  padding: 20px;
}

.header-row {
  display: flex;
  align-items: center;
  gap: 10px;
}

/* 移动端优化 */
@media (max-width: 768px) {
  /* 增加触摸反馈 */
  :deep(.el-table__row) {
    transition: background-color 0.2s;

    &:active {
      background-color: var(--el-fill-color-light);
    }
  }

  .el-header {
    padding: 10px;
  }

  .header-row {
    gap: 8px;
  }

  .search-input :deep(.el-input__wrapper) {
    border-radius: 20px;
  }

  .button-group .el-button {
    width: calc(50% - 4px);
    margin: 2px;
    padding: 8px;
  }

  .action-buttons {
    display: flex;
    flex-direction: column;
    gap: 4px;
  }

  .el-table {
    th,
    td {
      font-size: 12px;

      .cell {
        white-space: nowrap;
      }
    }
  }

  .el-dialog {
    border-radius: 8px;

    :deep(.el-dialog__body) {
      padding: 10px 15px;
    }
  }

  .el-form-item {
    margin-bottom: 12px;

    :deep(.el-form-item__label) {
      font-size: 13px;
    }
  }
}

@media (max-width: 480px) {
  .el-table-column--selection .cell {
    padding-left: 5px !important;
    padding-right: 5px !important;
  }

  .pagination :deep(.btn-prev),
  .pagination :deep(.btn-next) {
    min-width: 28px;
  }

  .pagination :deep(.number) {
    min-width: 28px;
  }
}
</style>
