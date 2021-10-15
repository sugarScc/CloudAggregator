<template>
  <div id="app">
    <el-container style="height: 100%;">
      <el-header class="el-header" >
        <!--引入header-->
        <home-header></home-header>
      </el-header>
      <el-container>
        <!--引入自定义左侧菜单栏-->
        <left-menu></left-menu>
        <el-main>
          <!--引入跳转路由-->
          <router-view></router-view>
        </el-main>
      </el-container>
    </el-container>
    <login :EnableLogin="EnableLogin"></login>
  </div>
</template>
<script>
import HomeHeader from './components/Header'
import LeftMenu from './components/LeftMenu'
import Login from './components/Login'

export default {
  name: 'App',
  data: function () {
    return {
      EnableLogin: false
    }
  },
  components: {
    HomeHeader,
    LeftMenu,
    Login
  },
  methods: {
  }
}
</script>

<style>
.el-header {
  background-color: #B3C0D1;
  color: #333;
  line-height: 60px;
}
#app {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
  height: 100%;
}
</style>
