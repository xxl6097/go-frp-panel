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
            <el-button type="warning" plain>导入用户</el-button>
          </template>
          <el-button
            type="warning"
            plain
            @click="handleExportUsers()"
            style="margin-left: 10px"
            >导出用户</el-button
          >
          <el-button
            type="primary"
            plain
            @click="showDialog('add', createEmptyUser())"
            >新增用户</el-button
          >
          <el-button type="success" plain @click="handleRefresh"
            >刷新</el-button
          >
          <el-popconfirm
            title="确定删除吗？"
            v-if="selectData.length !== 0"
            @confirm="handleDeleteUsers"
          >
            <template #reference>
              <el-button type="danger" plain>删除用户</el-button>
            </template>
          </el-popconfirm>
        </el-upload>
      </div>
    </el-header>

    <!-- 表格 -->
    <el-main>
      <el-table
        :data="paginatedTableData"
        style="width: 100%"
        @selection-change="handleSelectionChange"
        class="custom-border-table"
        :cell-style="{ padding: mobileLayout ? '4px' : '8px' }"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column prop="comment" label="备注" />
        <el-table-column prop="user" label="用户名" />
        <el-table-column prop="token" label="凭证" />
        <el-table-column prop="ports" label="允许端口" />
        <el-table-column prop="domains" label="允许域名" />
        <el-table-column prop="subdomains" label="允许子域名" />
        <el-table-column prop="enable" label="状态">
          <template #default="{ row }">
            <el-tag :type="row.enable ? 'success' : 'danger'">
              {{ row.enable ? '启动' : '禁用' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作">
          <template #default="{ row }">
            <el-dropdown
              :type="row.enable ? 'danger' : 'success'"
              size="small"
              placement="bottom"
              split-button
              plain
              @click="showDialog('ToggleStatus', row)"
            >
              {{ row.enable ? '禁用' : '启用' }}
              <template #dropdown>
                <el-dropdown-menu>
                  <el-dropdown-item @click="showDialog('update', row)"
                    >编辑
                  </el-dropdown-item>
                  <el-dropdown-item @click="handleDelete(row)"
                    >删除
                  </el-dropdown-item>
                  <el-dropdown-item @click="handleClientDialog(row)"
                    >生成客户端
                  </el-dropdown-item>
                </el-dropdown-menu>
              </template>
            </el-dropdown>
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
        <el-form-item label="用户名" prop="user">
          <el-input v-model="newUserForm.user" placeholder="请输入用户名(user)">
            <template #append>
              <el-button @click="handleRandUser">随机</el-button>
            </template>
          </el-input>
        </el-form-item>
        <el-form-item label="凭证" prop="token">
          <el-input
            v-model="newUserForm.token"
            placeholder="请输入Token凭证(meta_token)"
          />
        </el-form-item>
        <el-form-item label="备注">
          <el-input
            :rows="2"
            type="textarea"
            v-model="newUserForm.comment"
            placeholder="请输入备注"
          />
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
        <el-form-item label="服务器地址">
          <el-input v-model="clientForm.addr" placeholder="请输入服务器地址" />
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

        <el-form-item label="客户端下载地址" v-if="options.length <= 0">
          <el-input
            v-model="clientForm.url"
            placeholder="请输入客户端下载地址"
          />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="fetchClientToml">下载客户端toml配置</el-button>
        <el-button type="primary" @click="fetchClientGen()" :loading="isLoading"
          >确定</el-button
        >
      </template>
    </el-dialog>

    <!-- 删除 -->
  </el-container>
</template>

<script setup lang="ts">
import {
  ref,
  computed,
  onMounted,
  onUnmounted,
  reactive,
  onUpdated,
} from 'vue'
import {
  post,
  get,
  showErrorTips,
  showWarmDialog,
  generateRandomKey,
  deepCopyJSON,
  downloadByPost,
  showLoading,
} from '../utils/utils.ts'
import type { FormInstance, FormRules } from 'element-plus'

interface User {
  user: string
  token: string
  comment: string
  ports: string
  domains: string
  subdomains: string
  enable: boolean
  id: number
}

// const ops = ref({})
const options = ref([])
const isLoading = ref<boolean>(false)

// 搜索关键字
const searchKeyword = ref<string>('')
// 分页相关
const pageSize = ref<number>(10)
const currentPage = ref<number>(1)
//const filteredTableData = ref<User[]>([])

// 表格数据{user:'admin'}
//const tableData = ref<User[]>([{user: 'admin'}])
const tableData = ref<User[]>([])
const selectData = ref<User[]>([])
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
  id: 0,
})

const clientForm = ref({
  addr: '',
  url: '',
  ops: {},
})

const userRuleFormRef = ref<FormInstance>()

const userRules = reactive<FormRules>({
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

const resetForm = (formEl: FormInstance | undefined) => {
  if (!formEl) return
  formEl.resetFields()
}

const handleOptionChange = (value: any) => {
  console.log(value)
  //showSucessTips(JSON.stringify(node))
}
// 过滤后的表格数据（根据搜索关键字）
const filteredTableData = computed<User[]>(() => {
  return tableData.value.filter(
    (data) =>
      !searchKeyword.value ||
      data.user?.includes(searchKeyword.value) ||
      data.token?.includes(searchKeyword.value) ||
      data.comment?.includes(searchKeyword.value),
  )
})

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
    .then((data) => {
      console.log(data)
      //tableData.value = tableData.value.filter((item) => item !== row)
    })
    .catch((err) => {
      console.log(err)
    })
    .finally(() => {
      clearVariables()
      fetchData()
    })
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
const handleSelectionChange = (rows: User[]) => {
  selectData.value = rows
  console.log('--->', rows)
}

// 调用接口创建客户端
const fetchClientGen = () => {
  isLoading.value = genClientDialogVisible.value
  console.log('download----0----')
  console.log('fetchClientGen', clientForm.value.ops)
  const node = getFilePathByValue(options.value, clientForm.value.ops)
  const body = {
    user: newUserForm.value.user,
    token: newUserForm.value.token,
    comment: newUserForm.value.comment,
    ports: toPorts(newUserForm.value.ports),
    domains: newUserForm.value.domains.split(','),
    subdomains: newUserForm.value.subdomains.split(','),
    enable: newUserForm.value.enable,
  }
  if (node && node.filePath !== '') {
    //download('../api/client/gen?binPath=' + node.filePath)
    const data = {
      binPath: node.filePath,
      addr: clientForm.value.addr,
      user: body,
    }
    downloadByPost(
      '客户端生产中',
      '../api/client/gen',
      JSON.stringify(data),
    ).finally(() => {
      genClientDialogVisible.value = false
    })
    console.log('download----1----')
  } else {
    if (clientForm.value.url === '') {
      showErrorTips('生成客户端失败～')
      genClientDialogVisible.value = false
    } else {
      const data = {
        binUrl: clientForm.value.url,
        addr: clientForm.value.addr,
        user: body,
      }

      downloadByPost(
        '客户端生产中',
        '../api/client/gen',
        JSON.stringify(data),
      ).finally(() => {
        genClientDialogVisible.value = false
        isLoading.value = genClientDialogVisible.value
      })
      console.log('download----2----')
    }
  }
}

// 调用接口创建客户端
const fetchClientToml = () => {
  console.log('fetchClientToml', clientForm.value.ops)

  const body = {
    user: newUserForm.value.user,
    token: newUserForm.value.token,
    comment: newUserForm.value.comment,
    ports: toPorts(newUserForm.value.ports),
    domains: newUserForm.value.domains.split(','),
    subdomains: newUserForm.value.subdomains.split(','),
    enable: newUserForm.value.enable,
  }
  const data = {
    addr: clientForm.value.addr,
    user: body,
  }
  downloadByPost(
    '配置生成中',
    '../api/client/toml',
    JSON.stringify(data),
  ).finally(() => {
    genClientDialogVisible.value = false
  })
}

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
const paginatedTableData = computed<User[]>(() => {
  const start = (currentPage.value - 1) * pageSize.value
  const end = start + pageSize.value
  return filteredTableData.value.slice(start, end)
})

// 分页切换
const handlePageChange = (page: number) => {
  currentPage.value = page
}

const handleRefresh = () => {
  fetchData()
  fetchOptions()
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

const handleClientDialog = (row: User) => {
  genClientDialogVisible.value = true
  newUserForm.value = row
  console.log(row)
}

const showDialog = (type: string, row: User) => {
  clearVariables()

  //newUserForm.value = deepCopyJSON(row)
  //newUserForm.value = row
  if (type === 'ToggleStatus') {
    row.enable = !row.enable
    newUserForm.value = deepCopyJSON(row)
    updateUser()
  } else {
    newUserForm.value = deepCopyJSON(row)
    dialogVisible.value = true
    dialogType.value = type
  }
}

const handleDelete = (row: User) => {
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
        .then((data) => {
          console.log(data)
          tableData.value = tableData.value.filter((item) => item !== row)
        })
        .catch((err) => {
          console.log(err)
        })
        .finally(() => {
          clearVariables()
          fetchData()
        })
    },
    () => {
      clearVariables()
    },
  )
}

const handleRandUser = () => {
  newUserForm.value.token = `${generateRandomKey(16)}`
  newUserForm.value.user = `${new Date().getTime()}`
  console.log('handleRandUser', newUserForm.value)
}

const addUser = () => {
  const data = {
    user: newUserForm.value.user,
    token: newUserForm.value.token,
    comment: newUserForm.value.comment,
    ports: toPorts(newUserForm.value.ports),
    domains: newUserForm.value.domains.split(','),
    subdomains: newUserForm.value.subdomains.split(','),
    enable: newUserForm.value.enable,
  }
  const body = JSON.stringify(data)
  post('添加用户中...', '../api/token/add', body)
    .then((data) => {
      console.log(data)
      tableData.value.push({
        ...newUserForm.value,
        enable: true, // 默认状态为启用
      })
    })
    .catch((err) => {
      console.log(err)
    })
    .finally(() => {
      clearVariables()
      fetchData()
    })
}

const updateUser = () => {
  post('更新中...', '../api/token/chg', createUser(newUserForm.value))
    .then((data) => {
      console.log(data)
    })
    .catch((err) => {
      console.log(err)
    })
    .finally(() => {
      clearVariables()
      fetchData()
    })
}

const createUser = (row: User) => {
  const data = {
    user: row.user,
    token: row.token,
    comment: row.comment,
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
const createEmptyUser = () => {
  return {
    user: '',
    token: '',
    comment: '',
    ports: '',
    domains: '',
    subdomains: '',
    enable: true,
    id: 0,
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
  clientForm.value.addr = window.location.hostname
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

// 获取数据
const fetchData = () => {
  get('数据请求', '../api/token/all', null).then((data) => {
    if (data) {
      const obj = JSON.parse(JSON.stringify(data))
      tableData.value = obj.map((item: any) => ({
        user: item.user,
        token: item.token,
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
  get('', '../api/client/get', null).then((data) => {
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
fetchData()
fetchOptions()
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
