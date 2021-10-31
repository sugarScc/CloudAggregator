<template>
  <el-header style="text-align: right; font-size: 12px">
    <el-button v-if=" this.Global.AccountAddress=== undefined || this.Global.AccountAddress=== null" type="primary" @click="Login()">login</el-button>
    <el-button v-else type="primary" disabled>{{this.Global.AccountAddress}}</el-button>
  </el-header>
</template>

<script>
import Vue from 'vue'

export default {
  name: 'Header',
  data: function () {
  },
  methods: {
    Login: async function () {
      await this.getAccount()
    },
    async getAccount () {
      const accounts = await window.ethereum.request({method: 'eth_requestAccounts'})
      const account = accounts[0]
      this.Global.AccountAddress = account
    }
  }
}
</script>

<style scoped>
</style>
