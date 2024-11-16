<template>
  <el-config-provider :locale="ruRU">
    <div class=" bg-white px-8 py-2 border rounded-md flex flex-col gap-8 items-center text-left">
      <img :width="100" src="./assets/img.png" alt="Ugra logo"/>
      <div class="flex w-full flex-col items-start gap-4">
        <el-segmented class="w-full" v-model="value" :options="options"/>
      </div>
      <Login v-if="value === 'Вход'"/>
      <Register v-else/>
    </div>
  </el-config-provider>
</template>

<script setup lang="ts">
import ruRU from 'element-plus/es/locale/lang/ru'

import Login from "./components/Login.vue";
import {ref} from "vue";
import Register from "./components/Register.vue";

const value = ref('Вход')

const options = ['Вход', 'Регистрация']
</script>
