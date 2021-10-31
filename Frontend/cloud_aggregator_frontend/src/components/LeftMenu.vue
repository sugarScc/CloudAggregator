<template>
  <el-aside width="200px" style="background-color: rgb(238, 241, 246)">
    <el-menu router>
      <el-submeu index="1" >
        <template slot="title"><i class="el-icon-message"></i>导航一</template>
        <el-menu-item-group>
          <!--设置index-->
          <el-menu-item index="/taskForm">Create</el-menu-item>
          <el-menu-item index="/tasksInfo">TasksInfo</el-menu-item>
        </el-menu-item-group>
      </el-submeu>
    </el-menu>
  </el-aside>
</template>

<script>
export default {
  name: 'LeftMenu'
}
</script>

<style>
</style>
