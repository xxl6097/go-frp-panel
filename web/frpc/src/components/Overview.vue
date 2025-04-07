<template>
  <div>
    <el-page-header :icon="null" style="width: 100%; margin-bottom: 20px">
      <template #title>
        <el-select
          v-model="select_value"
          @change="handleSelectChange"
          placeholder="多客户端状态"
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
          <el-button type="success" :loading="loading" @click="refresh" plain
            >刷新</el-button
          >
        </div></template
      >
      <template #extra> </template>
    </el-page-header>

    <el-row>
      <el-col :md="24">
        <div>
          <el-table
            :data="status"
            stripe
            style="width: 100%"
            :default-sort="{ prop: 'type', order: 'ascending' }"
          >
            <el-table-column
              prop="name"
              label="名称"
              sortable
            ></el-table-column>
            <el-table-column
              prop="type"
              label="类型"
              width="150"
              sortable
            ></el-table-column>
            <el-table-column
              prop="local_addr"
              label="本地地址"
              width="200"
              sortable
            ></el-table-column>
            <el-table-column
              prop="plugin"
              label="插件"
              width="200"
              sortable
            ></el-table-column>
            <el-table-column
              prop="remote_addr"
              label="远程地址"
              sortable
            ></el-table-column>
            <el-table-column prop="status" label="状态" width="150" sortable>
              <template #default="{ row }">
                <el-tag :type="row.status === 'running' ? 'success' : 'danger'">
                  {{ row.status === 'running' ? '运行中' : '已停止' }}
                </el-tag>
              </template>
            </el-table-column>
            <el-table-column prop="err" label="信息"></el-table-column>
          </el-table>
        </div>
      </el-col>
    </el-row>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { ElMessage } from 'element-plus'
interface Option {
  value: string
  label: string
}
const status = ref<any[]>([])
const select_value = ref('')
const options = ref<Option[]>([])
const loading = ref<boolean>(false)

const handleSelectChange = (value: string) => {
  console.log('---->', value)
  if (value === '') {
    fetchData()
  } else {
    fetchStatus()
  }
}

const refresh = () => {
  loading.value = true
  console.log('refresh---->', select_value.value)
  if (select_value.value === undefined || select_value.value === '') {
    fetchData()
  } else {
    fetchStatus()
  }
}

const fetchListData = () => {
  fetch('../api/client/list', { credentials: 'include' })
    .then((res) => {
      return res.json()
    })
    .then((json) => {
      if (json.code === 0) {
        options.value = json.data
      }
    })
}

const fetchStatus = () => {
  const name = select_value.value
  fetch(`../api/client/status?name=${name}`, { credentials: 'include' })
    .then((res) => {
      return res.json()
    })
    .then((json) => {
      //status.value = new Array()
      status.value = []
      for (let key in json) {
        for (let ps of json[key]) {
          console.log(ps)
          status.value.push(ps)
        }
      }
    })
    .catch((err) => {
      ElMessage({
        showClose: true,
        message: 'Get status info from frpc failed!' + err,
        type: 'warning',
      })
    })
    .finally(() => {
      loading.value = false
    })
}

const fetchData = () => {
  fetch('/api/status', { credentials: 'include' })
    .then((res) => {
      return res.json()
    })
    .then((json) => {
      //status.value = new Array()
      status.value = []
      for (let key in json) {
        for (let ps of json[key]) {
          console.log(ps)
          status.value.push(ps)
        }
      }
    })
    .catch((err) => {
      ElMessage({
        showClose: true,
        message: 'Get status info from frpc failed!' + err,
        type: 'warning',
      })
    })
    .finally(() => {
      loading.value = false
    })
}
fetchData()
fetchListData()
</script>

<style></style>
