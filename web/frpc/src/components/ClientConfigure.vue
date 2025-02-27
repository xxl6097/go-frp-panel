<template>
  <div >
    <el-page-header
        :icon="null"
        style="width: 100%; margin-bottom: 20px"
    >
      <template #title>
        <el-select
            v-model="value"
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
        <div class="flex items-center" >
          <el-button type="primary" @click="upload" :loading="uploading" plain>更新</el-button>
          <el-button type="success" @click="refresh" :loading="loading" plain>刷新</el-button>
        </div></template>
      <template #extra>
      </template>
    </el-page-header>

    <el-input
      type="textarea"
      autosize
      style="margin-left: 10px"
      v-model="textarea"
      placeholder="frpc configure file, can not be empty..."
    ></el-input>
  </div>
</template>

<script setup lang="ts">
import { ref} from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import {showInfoTips} from "../utils/utils.ts";
interface Option{
  value: string
  label: string
}

const loading = ref<boolean>(false)
const uploading = ref<boolean>(false)
const textarea = ref('')
const value = ref('')
const options = ref<Option[]>([])
const handleSelectChange = (value:string) =>{
  console.log('---->',value)
  if (value === ''){
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
        if (json.code === 0){
          options.value = json.data
        }
      })
}
const fetchConfig = () =>{
  const name = value.value
  fetch(`../api/client/config/get?name=${name}`, { credentials: 'include' })
      .then((res) => {
        return res.text()
      })
      .then((text) => {
        if (text !== ''){
          textarea.value = text
        }
      }).finally(()=>{
        loading.value = false
  })
}
const fetchUpload = () =>{
  const data = {
    name: value.value,
    toml: textarea.value
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
        if (json.code === 0){
        }
        showInfoTips(json.msg)
      }).finally(()=>{
    uploading.value = false
  })
}
const refresh = () =>{
  loading.value = true
  if (value.value === ''){
    fetchData()
  }else{
    fetchConfig()
  }
}
const upload =()=>{
  uploading.value = true
  if (value.value === ''){
    uploadConfig()
  }else{
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
    }).finally(()=>{
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
    }
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
        }).finally(()=>{
        uploading.value = false
      })
    })
    .catch(() => {
      ElMessage({
        message: 'Canceled',
        type: 'info',
      })
    }).finally(()=>{
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
