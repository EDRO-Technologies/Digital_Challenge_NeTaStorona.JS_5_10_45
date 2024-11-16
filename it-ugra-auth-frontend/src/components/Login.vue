<template>
  <el-form v-loading="isFetching" class="w-[320px] flex flex-col" label-position="top" :model="formData">
    <el-form-item label="Email">
      <el-input size="large" v-model="formData.email" placeholder="Введите ваш email" clearable/>
    </el-form-item>
    <el-form-item label="Пароль">
      <el-input type="password" size="large" v-model="formData.password" placeholder="Введите пароль" clearable/>
    </el-form-item>
    <el-form-item>
      <el-button size="large" class="w-full mb-0 mt-6" type="primary" @click="onSubmit">Войти</el-button>
    </el-form-item>
  </el-form>
</template>

<script setup lang="ts">
import {reactive, ref} from 'vue'
import {ElNotification} from 'element-plus'
import axios from "axios";
import router from "../router";

function redirectToSite(jwtToken: string) {
  const baseUrl = import.meta.env.VITE_MAIN_FRONTEND_URL;

  const params = new URLSearchParams({
    token: jwtToken
  });

  window.location.href = `${baseUrl}?${params.toString()}`;
}

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

    localStorage.setItem('token', data.token);

    ElNotification.success({
      message: 'Приятного пользования сервисом!',
    })

    if (data.user.is_admin) {
      await router.push({path: '/agentreqs'})
    } else {
      redirectToSite(data.token)
    }
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