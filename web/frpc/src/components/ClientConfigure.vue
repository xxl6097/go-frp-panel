<template>
  <div>
    <el-page-header :icon="null" style="width: 100%; margin-bottom: 20px">
      <template #title>
        <el-select
          v-model="selectValue"
          @change="handleSelectChange"
          placeholder="请选择配置文件"
          clearable
          :fit-input-width="true"
          size="default"
          class="autoWidth1"
        >
          <el-option
            v-for="item in options"
            :key="item.value"
            :label="item.label"
            :value="item.value"
          />
        </el-select>
      </template>
      <template #content>
        <div class="flex items-center">
          <el-button type="primary" @click="upload" :loading="uploading" plain
            >更新</el-button
          >
          <el-button type="success" @click="refresh" :loading="loading" plain
            >刷新</el-button
          >
          <el-button
            type="info"
            @click="newClientFormVisible = true"
            :loading="loading"
            >新建客户端</el-button
          >

          <el-popconfirm title="确定删除客户端吗？" @confirm="deleteClient">
            <template #reference>
              <el-button
                type="danger"
                v-if="selectValue !== ''"
                :loading="loading"
                >删除客户端</el-button
              >
            </template>
          </el-popconfirm>
        </div></template
      >
      <template #extra> </template>
    </el-page-header>

    <el-input
      type="textarea"
      autosize
      style="margin-left: 10px"
      v-model="textarea"
      placeholder="frpc configure file, can not be empty..."
    ></el-input>
  </div>

  <!--新建客户端-->
  <el-dialog v-model="newClientFormVisible" title="创建客户端" width="700">
    <el-form ref="ruleFormRef" :model="newClientForm" :rules="rules">
      <el-form-item label="配置文件名：" prop="toml">
        <el-input
          v-model="newClientForm.name"
          placeholder="请输入toml配置文件名"
        />
      </el-form-item>

      <el-form-item prop="toml">
        <el-input
          type="textarea"
          rows="13"
          v-model="newClientForm.toml"
          placeholder="请在此输入toml格式配置内容"
        />
      </el-form-item>
    </el-form>
    <el-upload :http-request="uploadToml" :limit="1">
      <template #trigger>
        <el-link type="primary">上传toml配置文件</el-link>
      </template>
    </el-upload>
    <template #footer>
      <div class="dialog-footer">
        <el-button @click="newClientFormVisible = false">取消</el-button>
        <el-button type="primary" @click="submitForm(ruleFormRef)"
          >确定</el-button
        >
      </div>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">
import { reactive, ref } from 'vue'
import { ElMessage, ElMessageBox, FormInstance, FormRules } from 'element-plus'
import {
  put,
  showErrorTips,
  showInfoTips,
  showLoading,
  showSucessTips,
} from '../utils/utils.ts'
interface Option {
  value: string
  label: string
}

const newClientFormVisible = ref(false)
const loading = ref<boolean>(false)
const uploading = ref<boolean>(false)
const textarea = ref('')
const selectValue = ref('')
const options = ref<Option[]>([])

const newClientForm = ref({
  name: '',
  toml: '',
})

const ruleFormRef = ref<FormInstance>()
const rules = reactive<FormRules>({
  name: [
    {
      required: true,
      message: '请输入配置文件名',
      trigger: 'blur',
    },
  ],
  toml: [
    {
      required: true,
      message: '请输入配置内容',
      trigger: 'blur',
    },
  ],
})

const submitForm = async (formEl: FormInstance | undefined) => {
  if (!formEl) return
  await formEl.validate((valid, fields) => {
    if (valid) {
      console.log('submit!')
      handleNewFrpcClient()
    } else {
      console.log('error submit!', fields)
    }
  })
}

const handleNewFrpcClient = () => {
  const body = JSON.stringify(newClientForm.value)
  put('客户端创建中...', '../api/client/create', body).finally(() => {
    newClientFormVisible.value = false
    fetchListData()
  })
}

// 自定义上传函数
const uploadToml = (options: any) => {
  const { file } = options
  const formData = new FormData()
  formData.append('file', file)
  const loading = showLoading('客户端创建中...')
  // 使用 fetch 发送请求
  fetch('../api/client/create', {
    method: 'POST',
    body: formData,
  })
    .then((response) => {
      return response.json()
    })
    .then((data) => {
      // 上传成功的回调
      options.onSuccess(data)
    })
    .catch((error) => {
      // 上传失败的回调
      options.onError(error)
    })
    .finally(() => {
      loading.close()
      newClientFormVisible.value = false
      setTimeout(function () {
        window.location.reload()
      }, 1000)
    })
}

////
const handleSelectChange = (value: string) => {
  console.log('---->', value)
  if (value === '') {
    fetchData()
    return
  }
  fetchConfig()
}
const fetchListData = () => {
  fetch('../api/client/list', { credentials: 'include' })
    .then((res) => {
      return res.json()
    })
    .then((json) => {
      console.log('list', json)
      if (json.code === 0) {
        options.value = json.data
      }
    })
}
const fetchConfig = () => {
  const name = selectValue.value
  fetch(`../api/client/config/get?name=${name}`, { credentials: 'include' })
    .then((res) => {
      return res.json()
    })
    .then((json) => {
      if (json.code === 0) {
        textarea.value = json.data
      } else {
        showErrorTips(json.msg)
      }
    })
    .finally(() => {
      loading.value = false
    })
}
const fetchUpload = () => {
  const data = {
    name: selectValue.value,
    toml: textarea.value,
  }
  const body = JSON.stringify(data)
  fetch(`../api/client/config/set`, {
    credentials: 'include',
    method: 'POST',
    body: body,
  })
    .then((res) => {
      return res.json()
    })
    .then((json) => {
      showInfoTips(json.msg)
    })
    .finally(() => {
      uploading.value = false
    })
}
const refresh = () => {
  loading.value = true
  if (selectValue.value === '') {
    fetchData()
  } else {
    fetchConfig()
  }
  fetchListData()
}

const deleteClient = () => {
  if (selectValue.value !== '') {
    loading.value = true
    fetch(`../api/client/delete?name=${selectValue.value}`, {
      credentials: 'include',
      method: 'DELETE',
    })
      .then((res) => {
        return res.json()
      })
      .then((json) => {
        if (json.code === 0) showSucessTips(json.msg)
      })
      .catch(() => {
        ElMessage({
          showClose: true,
          message: 'delete frpc failed!',
          type: 'warning',
        })
      })
      .finally(() => {
        loading.value = false
        selectValue.value = ''
        fetchListData()
        fetchData()
      })
  }
}
const upload = () => {
  uploading.value = true
  if (selectValue.value === '') {
    uploadConfig()
  } else {
    fetchUpload()
  }
}

const fetchData = () => {
  fetch('/api/config', { credentials: 'include' })
    .then((res) => {
      return res.text()
    })
    .then((text) => {
      textarea.value = text
    })
    .catch(() => {
      ElMessage({
        showClose: true,
        message: 'Get configure content from frpc failed!',
        type: 'warning',
      })
    })
    .finally(() => {
      loading.value = false
    })
}
const uploadConfig = () => {
  ElMessageBox.confirm(
    'This operation will upload your frpc configure file content and hot reload it, do you want to continue?',
    'Notice',
    {
      confirmButtonText: 'Yes',
      cancelButtonText: 'No',
      type: 'warning',
    },
  )
    .then(() => {
      if (textarea.value == '') {
        ElMessage({
          message: 'Configure content can not be empty!',
          type: 'warning',
        })
        return
      }

      fetch('/api/config', {
        credentials: 'include',
        method: 'PUT',
        body: textarea.value,
      })
        .then(() => {
          fetch('/api/reload', { credentials: 'include' })
            .then(() => {
              ElMessage({
                type: 'success',
                message: 'Success',
              })
            })
            .catch((err) => {
              ElMessage({
                showClose: true,
                message: 'Reload frpc configure file error, ' + err,
                type: 'warning',
              })
            })
        })
        .catch(() => {
          ElMessage({
            showClose: true,
            message: 'Put config to frpc and hot reload failed!',
            type: 'warning',
          })
        })
        .finally(() => {
          uploading.value = false
        })
    })
    .catch(() => {
      ElMessage({
        message: 'Canceled',
        type: 'info',
      })
    })
    .finally(() => {
      uploading.value = false
    })
}

fetchData()
fetchListData()
</script>

<style>
#head {
  margin-bottom: 30px;
}
.autoWidth1 {
  width: auto;
  min-width: 250px; /* 初始最小宽度 */
  max-width: 400px; /* 初始最小宽度 */
  margin-left: 10px;
}
</style>
