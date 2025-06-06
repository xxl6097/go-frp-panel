<template>
  <el-container>
    <!-- 表格 -->
    <el-main>
      <el-table
        :data="paginatedTableData"
        style="width: 100%"
        :border="true"
        :preserve-expanded-content="true"
        :cell-style="{ padding: mobileLayout ? '4px' : '8px' }"
      >
        <el-table-column type="expand">
          <template #default="props">
            <div m="4">
              <p m="t-0 b-2">设备名称: {{ props.row.devName }}</p>
              <p m="t-0 b-2">版本号: {{ props.row.appVersion }}</p>
              <p m="t-0 b-2">允许端口: {{ frpConfig?.ports }}</p>
              <p m="t-0 b-2">Frp连接ID: {{ props.row.frpId }}</p>
              <p m="t-0 b-2">操作系统: {{ props.row.osType }}</p>
              <p m="t-0 b-2">设备Mac地址: {{ props.row.devMac }}</p>
              <p m="t-0 b-2">设备IP地址: {{ props.row.devIp }}</p>
              <p m="t-0 b-2">客户端websocket连接ID: {{ props.row.secKey }}</p>
            </div>
          </template>
        </el-table-column>
        <el-table-column prop="devName" label="设备名称" />
        <el-table-column prop="frpId" label="Frp连接ID" />
        <el-table-column prop="appVersion" label="版本号" />
        <el-table-column prop="devMac" label="设备Mac" />
        <el-table-column label="操作">
          <template #default="{ row }">
            <el-button
              size="small"
              type="success"
              plain
              @click="handleGoToDetail(row)"
              >查看
            </el-button>
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
  </el-container>

  <ClientDetailDialog ref="clientDetailDialogRef" />
</template>

<script setup lang="ts">
import { ref, computed, onMounted, onUnmounted, onUpdated } from 'vue'
import ClientDetailDialog from './ClientDetailDialog.vue'
import { Client, FrpcConfiguration } from '../../utils/type.ts'
import { useRoute } from 'vue-router'

// 搜索关键字
const searchKeyword = ref<string>('')
// 分页相关
const pageSize = ref<number>(10)
const currentPage = ref<number>(1)
const tableData = ref<Client[]>([])
const frpConfig = ref<FrpcConfiguration>()

const clientForm = ref({
  addr: '',
  port: 0,
  url: '',
  ops: {},
})

const clientDetailDialogRef = ref<InstanceType<
  typeof ClientDetailDialog
> | null>(null)

// 过滤后的表格数据（根据搜索关键字）
const filteredTableData = computed<Client[]>(() => {
  return tableData.value.filter(
    (data) =>
      !searchKeyword.value || data.secKey?.includes(searchKeyword.value),
  )
})

// 分页后的表格数
const paginatedTableData = computed<Client[]>(() => {
  const start = (currentPage.value - 1) * pageSize.value
  const end = start + pageSize.value
  return filteredTableData.value.slice(start, end)
})

// 分页切换
const handlePageChange = (page: number) => {
  currentPage.value = page
}

// 调整详情
const handleGoToDetail = (row: Client) => {
  console.log('handleGoToDetail', row)
  if (clientDetailDialogRef.value) {
    clientDetailDialogRef.value.openClientDialog(row, frpConfig.value)
  }
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

const fetchListData = (id: string) => {
  const data = {
    frpId: id,
  }
  console.log('fetchListData.query', data)
  fetch('../api/client/list', {
    credentials: 'include',
    method: 'POST',
    body: JSON.stringify(data),
  })
    .then((res) => res.json())
    .then((json) => {
      console.log('fetchListData.json', json)
      if (json.code === 0) {
        tableData.value = json.data.map((item: any) => ({
          osType: item.osType,
          secKey: item.secKey,
          devMac: item.devMac,
          devIp: item.devIp,
          frpId: item.frpId,
          appVersion: item.appVersion,
          devName: item.devName,
        }))
      }
    })
    .catch(() => {
      // showErrorTips('获取服务器信息失败')
    })
}

const fetchFrpcData = (id: string) => {
  console.log('fetchFrpcData', id)
  fetch(`../api/token/get?id=${id}`, {
    credentials: 'include',
    method: 'GET',
  })
    .then((res) => res.json())
    .then((json) => {
      console.log('fetchFrpcData', json)
      if (json.code === 0) {
        const cfg = {
          user: json.data.user,
          token: json.data.token,
          count: json.data.count,
          comment: json.data.comment,
          ports: json.data.ports.join(','),
          domains: json.data.domains.join(','),
          subdomains: json.data.subdomains.join(','),
          enable: json.data.enable,
          id: json.data.id,
        }
        frpConfig.value = cfg as FrpcConfiguration
        console.log('frpConfig', frpConfig.value)
      }
    })
    .catch(() => {
      // showErrorTips('获取服务器信息失败')
    })
}

// const props = defineProps<{
//   profile: FrpcConfiguration
// }>()
onUpdated(() => {})
const id = useRoute().query.id as string
fetchFrpcData(id)
fetchListData(id)
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
