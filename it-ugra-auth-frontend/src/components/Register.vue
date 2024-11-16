<template>
  <el-form class="w-[320px] flex flex-col " label-width="auto" label-position="top" :model="formData">
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
      <el-input size="large" v-model="formData.password" placeholder="Иванов Иван Иванович" clearable/>
    </el-form-item>
    <el-form-item label="Дата рождения">
      <el-date-picker
          v-model="formData.birthdate"
          size="large"
          type="date"
          placeholder="Когда у Вас день рождения?"
          clearable
      />
    </el-form-item>



    <el-form-item>
      <el-button size="large" class="w-full mb-0" type="primary" @click="onSubmit">Зарегистрироваться</el-button>
    </el-form-item>
  </el-form>
</template>

<script setup lang="ts">
import {reactive} from 'vue'

const formData = reactive({
  email: '',
  password: '',
  passwordConfirmation: '',
  birthdate: new Date()
})

const onSubmit = () => {
  console.log({formData})
}
</script>

<style scoped>

</style>