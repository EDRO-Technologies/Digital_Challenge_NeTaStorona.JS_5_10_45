<template>
  <el-form v-loading="isFetching" class="w-[320px] flex flex-col" label-position="top" :model="formData">
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
import {reactive, ref} from 'vue'
import {ElNotification} from 'element-plus'
import axios from "axios";

const isFetching = ref(false);

const formData = reactive({
  email: '',
  password: '',
})

const onSubmit = async () => {
  const url = import.meta.env.VITE_AUTH_SERVER_URL + '/auth/login'

  isFetching.value = true;

  try {
    const {data} = await axios.post(url, {
      email: formData.email,
      password: formData.password,
    })

    ElNotification.success({
      message: 'Приятного пользования сервисом!',
    })
    console.log({data})
  } catch (error) {
    console.error(error)
    ElNotification.error({
      message: 'Неверный логин или пароль :(',
    })
  }

  isFetching.value = false;
}
</script>

<style scoped>

</style>