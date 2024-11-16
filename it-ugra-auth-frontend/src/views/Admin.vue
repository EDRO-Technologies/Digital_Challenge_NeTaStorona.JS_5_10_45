<template>

  <div v-if="agentRequestList.length !== 0" v-loading="isFetching" class="flex flex-col w-full max-w-[1280px] gap-4">
    <AgentRequest v-for="request in agentRequestList" :request="request" />
  </div>
  <div v-else class="bg-white/60 px-4 py-4 rounded-md border w-64">
    <el-empty class="text-black"/>
  </div>



</template>

<script setup lang="ts">
import {onBeforeMount, ref} from "vue";
import axios from "axios";
import {ElNotification} from "element-plus";
import {AgentRequestType} from "../types/AgentRequestType.ts";

import AgentRequest from "../components/AgentRequest.vue";

const isFetching = ref(false);

const agentRequestList = ref<AgentRequestType[]>([]);

onBeforeMount(async () => {
  const url = import.meta.env.VITE_AUTH_SERVER_URL + '/auth/agent/requests'

  isFetching.value = true;

  try {
    const token = localStorage.getItem('token');

    const {data} = await axios.get<AgentRequestType[]>(url, {headers: {"Authorization": "Bearer " + token}});

    agentRequestList.value = data;

    ElNotification.success({
      message: 'Приятного пользования сервисом!',
    })
  } catch (error) {
    console.error(error)
    ElNotification.error({
      message: 'У вас нет прав для просмотра данной страницы',
    })
  }

  isFetching.value = false;
})
</script>
