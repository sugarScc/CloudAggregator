<template>

  <el-table :data="tableData">
    <el-table-column prop="transaction_id" label="Id" width="120">
    </el-table-column>
    <el-table-column prop="state" label="Status" width="140">
      <template slot-scope="scope">
        <el-result v-if="scope.row.state==='Waiting'" icon="warning"></el-result>
        <el-result v-if="scope.row.state==='Succeeded'" icon="success"></el-result>
        <el-result v-if="scope.row.state==='Failed'" icon="error"></el-result>
        <el-result v-if="scope.row.state==='Canceled'" icon="info"></el-result>
      </template>
    </el-table-column>
    <el-table-column prop="creation_time_stamps" label="CreationTime" width="120">
    </el-table-column>
    <el-table-column prop="docker_image" label="DockerImage" width="200">
    </el-table-column>
    <el-table-column prop="flag_message" label="LivenessMsg" width="120">
    </el-table-column>
    <el-table-column prop="port" label="LivenessPort" width="120">
    </el-table-column>
    <el-table-column prop="url" label="URL" width="120">
    </el-table-column>
    <el-table-column prop="with_draw_visible" label="withdraw" width="140">
      <template scope="scope">
        <el-button type="success" v-if="scope.row.with_draw_visible" icon="error" @click="withdraw()">withdraw
        </el-button>
      </template>
    </el-table-column>
  </el-table>
</template>

<script>
import axios from 'axios'

export default {
  name: 'TasksInfo',
  data () {
    return {
      tableData: null
    }
  },
  mounted () {
    axios.get('/api/v1/transactions/user/' + this.Global.AccountAddress).then(
      response => {
        this.tableData = []
        for (let i = 0; i < response.data.Data.length; i++) {
          console.log(response.data.Data[i])
          this.tableData.push(response.data.Data[i])
        }
        console.log(this.tableData)
      }
    )
  },
  methods: {
    withdraw () {
      axios.get('/api/v1/transactions/user/' + this.Global.AccountAddress).then(
        response => {
          this.tableData = []
          for (let i = 0; i < response.data.Data.length; i++) {
            console.log(response.data.Data[i])
            this.tableData.push(response.data.Data[i])
          }
          console.log(this.tableData)
        }
      )
    }
  }
}
</script>

<style scoped>

</style>
