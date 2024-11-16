<template>
  <el-container class="gap-14">
    <el-header>
      <Header />
    </el-header>
    <el-main>
      <router-view />
    </el-main>
    <el-footer></el-footer>
  </el-container>
</template>

<script setup lang="ts">
import Header from "@/views/components/Header.vue";
</script>
