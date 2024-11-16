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
import { onMounted } from "vue";
import { sendLlmRequest } from "@/services/LegoService";
import type { ILegoRequest } from "@/stores/models/LegoDto";

const request = {
  context:
    "Я участвовал в : 1. Хантатон. Бэкенд разработчик на C; 2. Хакатон. Бэкенд разработчик на Python; Доступны мероприятия: 1. Соревнования по бегу; 2. Хоккейный матч; 3. Хакатон; 4. Соревнования по готовке; 5. Конкурс программирования; 6. IT-challenge; 7. Плавание; 8. Хантатон; 9. Конкурс весёлых историй",
  message:
    "Перечисли мероприятия, в которых мне желательней всего участвовать? Дай ответ исключительно в следующем виде: id,id,..",
  user_id: 0,
} as ILegoRequest;

onMounted(async () => {
  await sendLlmRequest(request);
});
</script>
