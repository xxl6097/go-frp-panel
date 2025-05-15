<template>
  <el-container>
    <!-- 表格 -->
    <el-main>
      <el-table
        :data="paginatedTableData"
        style="width: 100%"
        class="custom-border-table"
        :cell-style="{ padding: mobileLayout ? '4px' : '8px' }"
      >
        <el-table-column prop="frpId" label="FrpID" />
        <el-table-column
          prop="secKey"
          label="会话ID"
          width="150"
          :show-overflow-tooltip="true"
        />
        <el-table-column prop="osType" label="操作系统" />
        <el-table-column prop="devMac" label="设备Mac" />
        <el-table-column prop="devIp" label="设备IP" />
        <el-table-column label="操作">
          <template #default="{ row }">
            <el-button size="small" @click="handleGoToDetail(row)"
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
import { useRoute } from 'vue-router'
import ClientDetailDialog from './client/ClientDetailDialog.vue'
import { Client } from '../utils/type.ts'

// 搜索关键字
const searchKeyword = ref<string>('')
// 分页相关
const pageSize = ref<number>(10)
const currentPage = ref<number>(1)
const tableData = ref<Client[]>([])
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
    clientDetailDialogRef.value.openClientDialog(row)
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

const fetchListData = () => {
  const data = {
    frpId: useRoute().query.frpId,
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
        }))
      }
    })
    .catch(() => {
      // showErrorTips('获取服务器信息失败')
    })
}
// 获取数据
// const fetchListData = () => {
//   const data = {
//     timeId: query.timeId,
//   }
//   get('数据请求', '../api/client/list', JSON.stringify(data)).then(
//     (data: any) => {
//       console.log('fetchListData', data)
//       if (data) {
//         const obj = JSON.parse(JSON.stringify(data))
//         tableData.value = obj.map((item: any) => ({
//           osType: item.osType,
//           secKey: item.secKey,
//           devMac: item.devMac,
//           devIp: item.devIp,
//           id: item.id,
//         }))
//       } else {
//         tableData.value = []
//       }
//     },
//   )
// }

onUpdated(() => {})
onMounted(() => {})
fetchListData()
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
