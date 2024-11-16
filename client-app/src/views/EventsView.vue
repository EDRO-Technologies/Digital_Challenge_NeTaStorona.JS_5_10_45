<template>
  <div class="flex items-center justify-center gap-2">
    <el-date-picker
      v-model="dateRangeValue"
      type="daterange"
      range-separator="По"
      class="max-w-sm min-w-[36vw]"
    />
    <el-select v-model="filterValue" placeholder="Выберите фильтр">
      <el-option
        v-for="item in options"
        :key="item.value"
        :label="item.label"
        :value="item.value"
      />
    </el-select>
    <el-input
      v-model="searchValue"
      :prefix-icon="Search"
      placeholder="Поиск..."
    />
  </div>
  <el-tabs v-model="eventType">
    <el-tab-pane label="Хакатоны" name="hack">
      <EventCard title="Мероприятие 1" description="" source="/Pattern.png" />
      <EventCard title="SurGu Digital Challenge" description="" source="/Pattern2.png" />
      <EventCard title="Еще мероприятие" description="" source="/Pattern3.png" />
    </el-tab-pane>
    <el-tab-pane label="Акселераторы" name="accelerator">
      <el-empty description="Здесь будут отображаться мероприятия" />
    </el-tab-pane>
    <el-tab-pane label="Форумы" name="forum">
      <el-empty description="Здесь будут отображаться мероприятия" />
    </el-tab-pane>
    <el-tab-pane label="Лекции" name="lecture" class="flex flex-col gap-2">
      <el-empty description="Здесь будут отображаться мероприятия" />
    </el-tab-pane>
  </el-tabs>
</template>

<script setup lang="ts">
import { ref } from "vue";
import { Search } from "@element-plus/icons-vue";
import EventCard from "@/views/components/EventCard.vue";

const eventType = ref("hack");
const filterValue = ref("");
const dateRangeValue = ref("");
const searchValue = ref("");

const options = [
  {
    value: "Option1",
    label: "По дате создания",
  },
  {
    value: "Option2",
    label: "По актуальности",
  },
  {
    value: "Option3",
    label: "По другому параметру",
  },
  {
    value: "Option4",
    label: "Еще по другому параметру",
  },
  {
    value: "Option5",
    label: "Еще по другому другому параметру",
  },
];
</script>

<style scoped></style>
