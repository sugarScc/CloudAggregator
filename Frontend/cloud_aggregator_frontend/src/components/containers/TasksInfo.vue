<template>
  <el-table :data="tableData">
    <el-table-column prop="status" label="部署情况" width="140">
      <template scope="scope">
        <el-result v-if="scope.row.status==='success'" icon="success"></el-result>
        <el-result v-if="scope.row.status==='failed'" icon="error"></el-result>
      </template>
    </el-table-column>
    <el-table-column prop="createTime" label="创建时间" width="120">
    </el-table-column>
    <el-table-column prop="deployedTime" label="部署时间" width="120">
    </el-table-column>
    <el-table-column prop="ip" label="检验ip" width="120">
    </el-table-column>
    <el-table-column prop="provider" label="服务提供商" width="120">
    </el-table-column>
    <el-table-column prop="retrieveDeposit" label="取回定金" width="120">
      <el-button type="warning">test</el-button>
    </el-table-column>
  </el-table>
</template>

<script>
export default {
  name: 'TasksInfo',
  data () {
    const item = {
      status: 'failed',
      createTime: '2021-05-02',
      deployedTime: '2021-05-02',
      ip: '127.0.0.1',
      provider: 'CUHKCloud',
    }
    return {
      tableData: Array(20).fill(item)
    }
  }
}
</script>

<style scoped>

</style>
