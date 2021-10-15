<template>
  <el-header style="text-align: right; font-size: 12px">
    <el-button type="primary" @click="Login()">登录</el-button>
  </el-header>
</template>

<script>
import {eventBus} from '../eventBus'

export default {
  name: 'Header',
  methods: {
    Login: function () {
      console.log('true')
      eventBus.$emit('login')
    }

  }
}
</script>

<style scoped>
</style>
