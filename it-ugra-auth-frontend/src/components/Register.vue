<template>
  <el-form v-loading="isFetching" class="w-[320px] flex flex-col " label-width="auto" label-position="top" :model="formData">
    <el-form-item label="Email">
      <el-input size="large" v-model="formData.email" placeholder="Введите ваш email" clearable/>
    </el-form-item>
    <el-form-item label="Пароль">
      <el-input type="password" size="large" v-model="formData.password" placeholder="Введите пароль" clearable/>
    </el-form-item>
    <el-form-item label="Повтор пароля">
      <el-input type="password" size="large" v-model="formData.passwordConfirmation" placeholder="Повторите пароль" clearable/>
    </el-form-item>
    <el-form-item label="ФИО">
      <el-input size="large" v-model="formData.fio" placeholder="Иванов Иван Иванович" clearable/>
    </el-form-item>
    <el-form-item label="Дата рождения">
      <el-date-picker
          v-model="formData.birthdate"
          size="large"
          type="date"
          placeholder="Когда у Вас день рождения?"
          format="DD.MM.YYYY"
          clearable
      />
    </el-form-item>

    <el-form-item>
      <el-button size="large" class="w-full mb-0 mt-6" type="primary" @click="onSubmit">Зарегистрироваться</el-button>
    </el-form-item>
  </el-form>
</template>

<script setup lang="ts">
import {reactive, ref} from 'vue'
import axios from "axios";
import {ElNotification} from "element-plus";

const formData = reactive({
  email: '',
  password: '',
  passwordConfirmation: '',
  birthdate: new Date(),
  fio: ''
})

const isFetching = ref(false);

const onSubmit = async () => {
  const url = import.meta.env.VITE_AUTH_SERVER_URL + '/auth/register'

  isFetching.value = true;

  if (formData.password !== formData.passwordConfirmation) {
    ElNotification.error({
      message: 'Пароли не совпадают!',
    })

    isFetching.value = false;
    return;
  }

  try {
    const {data} = await axios.post(url, {
      email: formData.email,
      password: formData.password,
      birthdate: new Date(),
      fio: formData.fio,
    })

    ElNotification.success({
      message: 'Приятного пользования сервисом!',
    })
    console.log({data})
  } catch (error: any) {
    console.error(error)

    if (error?.response?.status === 403 || error?.response?.status === 400) {
      ElNotification.error({message: error.response.data.msg})
    } else {
      ElNotification.error({
        message: 'Что то пошло не так(((',
      })
    }
  }

  isFetching.value = false;
}
</script>

<style scoped>

</style>