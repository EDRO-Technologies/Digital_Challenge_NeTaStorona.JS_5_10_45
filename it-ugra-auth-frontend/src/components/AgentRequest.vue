
<template>
  <div v-loading="isApproving" class="min-w-full w-full text-black text-left px-4 py-2 bg-white rounded-md border-2 flex flex-col gap-4" >

    <div class="flex flex-row justify-between w-full text-xl border-b-2 pb-2 font-bold">
      <div>{{props.request.fio}}</div>
      <div>{{props.request.email}}</div>
    </div>

    <p v-html="request.agent_request"/>

    <el-button v-if="!isApproved" type="primary" @click="onApprove">Сделать интересодержателем</el-button>
    <el-button v-else disabled type="success">{{props.request.fio}} теперь интересодержатель</el-button>
  </div>
</template>

<script setup lang="ts">
import {AgentRequestType} from "../types/AgentRequestType.ts";
import {ref} from "vue";
import axios from "axios";
import {ElNotification} from "element-plus";

const props = defineProps<{
  request: AgentRequestType,
}>()

const isApproving = ref(false);
const isApproved = ref(false);

async function onApprove() {
  const url = import.meta.env.VITE_AUTH_SERVER_URL + `/auth/agent/requests/${props.request.id}/approve`

  isApproving.value = true;

  try {
    const token = localStorage.getItem('token');

    await axios.post<AgentRequestType[]>(url, undefined, {headers: {"Authorization": "Bearer " + token}});

    isApproved.value = true;

    ElNotification.success({
      message: `${props.request.fio} теперь интересодержатель`,
    })
  } catch (error) {
    console.error(error)
    ElNotification.error({
      message: 'Что-то пошло не так :(',
    })
  }

  isApproving.value = false;
}

</script>
