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
            >刷新
          </el-button>
        </div>
      </template>
      <template #extra></template>
    </el-page-header>

    <el-row>
      <el-col :md="24">
        <div>
          <el-table
            :data="frpcProxyList"
            stripe
            style="width: 100%"
            :default-sort="{ prop: 'type', order: 'ascending' }"
          >
            <!-- 手机端：展开查看完整详情 -->
            <el-table-column v-if="isMobile" type="expand">
              <template #default="{ row }">
                <div class="proxy-expand">
                  <div><span class="lbl">类型</span>{{ row.type }}</div>
                  <div>
                    <span class="lbl">本地地址</span>{{ row.local_addr || '—' }}
                  </div>
                  <div>
                    <span class="lbl">插件</span>{{ row.plugin || '—' }}
                  </div>
                  <div>
                    <span class="lbl">远程地址</span>{{ row.remote_addr || '—' }}
                  </div>
                  <div><span class="lbl">信息</span>{{ row.err || '—' }}</div>
                </div>
              </template>
            </el-table-column>
            <el-table-column
              prop="name"
              label="名称"
              min-width="120"
              sortable
              show-overflow-tooltip
            ></el-table-column>
            <el-table-column
              v-if="!isMobile"
              prop="type"
              label="类型"
              width="150"
              sortable
            ></el-table-column>
            <el-table-column
              v-if="!isMobile"
              prop="local_addr"
              label="本地地址"
              width="200"
              sortable
            ></el-table-column>
            <el-table-column
              v-if="!isMobile"
              prop="plugin"
              label="插件"
              width="200"
              sortable
            ></el-table-column>
            <el-table-column
              v-if="!isMobile"
              prop="remote_addr"
              label="远程地址"
              min-width="140"
              sortable
              show-overflow-tooltip
            ></el-table-column>
            <el-table-column
              prop="status"
              label="状态"
              :width="isMobile ? 90 : 150"
              sortable
            >
              <template #default="{ row }">
                <el-tag
                  :type="row.status === 'running' ? 'success' : 'danger'"
                  size="small"
                >
                  {{ row.status === 'running' ? '运行中' : '已停止' }}
                </el-tag>
              </template>
            </el-table-column>
            <el-table-column
              v-if="!isMobile"
              prop="err"
              label="信息"
              min-width="120"
              show-overflow-tooltip
            ></el-table-column>
          </el-table>
        </div>
      </el-col>
    </el-row>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { ElMessage } from 'element-plus'
import { showWarmDialog } from '../utils/utils.ts'
import { useResponsive } from '../composables/useResponsive'

const { isMobile } = useResponsive()

interface Option {
  value: string
  label: string
}

const frpcProxyList = ref<any[]>([])
const select_value = ref('')
const options = ref<Option[]>([])
const loading = ref<boolean>(false)

const handleSelectChange = (value: string) => {
  console.log('---->', value)
  if (value === undefined || value === '') {
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
      console.log('fetchListData api/client/list', json)
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
      console.log('fetchStatus', json)
      //status.value = new Array()
      // status.value = []
      if (json.code === 0 && json.data) {
        frpcProxyList.value.splice(0, frpcProxyList.value.length)
        const data = json.data
        for (let key in data) {
          for (let ps of data[key]) {
            //console.log(ps)
            frpcProxyList.value.push(ps)
          }
        }
      } else {
        // showTips(json.code, json.msg)
        showWarmDialog(json.msg, {}, {})
      }
    })
    .catch((err) => {
      console.error('fetchStatus err', err)
      ElMessage({
        showClose: true,
        message: '从 frpc 获取状态信息失败！' + err,
        type: 'warning',
      })
    })
    .finally(() => {
      loading.value = false
    })
}

const fetchData = () => {
  fetch('../api/status', { credentials: 'include' })
    .then((res) => {
      return res.json()
    })
    .then((json) => {
      console.log('fetchData /api/status', json)
      //status.value = new Array()
      // status.value = []
      frpcProxyList.value.splice(0, frpcProxyList.value.length)
      for (let key in json) {
        for (let ps of json[key]) {
          //console.log(ps)
          frpcProxyList.value.push(ps)
        }
      }
    })
    .catch((err) => {
      console.error('fetchData', err)
      ElMessage({
        showClose: true,
        message: '从 frpc 获取状态信息失败！' + err,
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

<style scoped>
/* 手机端展开详情 */
.proxy-expand {
  padding: 8px 16px;
  display: flex;
  flex-direction: column;
  gap: 6px;
  font-size: 13px;
}
.proxy-expand .lbl {
  display: inline-block;
  min-width: 72px;
  color: var(--el-text-color-secondary);
  font-weight: 500;
}
</style>
