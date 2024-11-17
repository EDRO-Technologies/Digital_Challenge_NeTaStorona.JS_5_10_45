<template>
  <div class="flex h-[26vh]">
    <div class="date-card">
      <p>{{ dayjs(props.event.date_start).format("DD") }}</p>
      <p>{{ dayjs(props.event.date_start).format("MMM") }}</p>
    </div>
    <img
      :src="props.event.photo"
      alt="Изображение мероприятия"
      width="200"
      height="200"
    />
    <div class="flex flex-col gap-1 ml-2">
      <p class="font-bold text-xl">{{ props.event.full_name }}</p>
      <div class="h-15px overflow-y-hidden">
        <p v-html="props.event.full_info"></p>
      </div>
      <el-button type="primary" class="self-end max-w-40"
        >Подробнее
      </el-button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { AppEvent } from "@/stores/models/AppEvent";
import dayjs from "dayjs";

const props = defineProps<{
  event: AppEvent;
}>();
</script>

<style scoped>
.date-card {
  @apply w-[70px] font-bold text-xl flex flex-col items-start justify-center bg-cyan-800 text-white px-2 -mr-0.5;
}
</style>
