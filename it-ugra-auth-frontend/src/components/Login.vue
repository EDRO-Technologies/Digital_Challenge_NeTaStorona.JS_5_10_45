<template>
  <el-form class="w-[320px] flex flex-col" label-position="top" :model="formData">
    <el-form-item label="Email">
      <el-input size="large" v-model="formData.email" placeholder="Введите ваш email" clearable/>
    </el-form-item>
    <el-form-item label="Пароль">
      <el-input type="password" size="large" v-model="formData.password" placeholder="Введите пароль" clearable/>
    </el-form-item>
    <el-form-item>
      <el-button size="large" class="w-full mb-0" type="primary" @click="onSubmit">Войти</el-button>
    </el-form-item>
  </el-form>
</template>

<script setup lang="ts">
import {reactive} from 'vue'

const formData = reactive({
  email: '',
  password: '',
})

const onSubmit = () => {
  console.log({formData})
}
</script>

<style scoped>

</style>