<template>
  <div class="header">
    <el-menu
      mode="horizontal"
      :router="true"
      >
      <el-menu-item index="1" route="/home">Home</el-menu-item>
      <el-menu-item index="2" route="/tutorials">Tutorials</el-menu-item>
      <el-menu-item index="3" route="/users/login">Login</el-menu-item>
      
    </el-menu>
  </div>
</template>
