<template>
  <div
    v-loading.fullscreen.lock="isLoading"
    class="flex items-center justify-center gap-2"
  >
    <el-date-picker
      v-model="dateRangeFilter"
      type="daterange"
      range-separator="По"
      class="max-w-sm min-w-[36vw]"
    />
    <el-select v-model="typeFilter" placeholder="Выберите фильтр" clearable>
      <el-option
        v-for="eventType in AppEventTypesEnum"
        :key="eventType"
        :label="eventType"
        :value="eventType"
      />
    </el-select>
    <el-input
      v-model="nameSearchFilter"
      :prefix-icon="Search"
      placeholder="Поиск..."
    />
  </div>
  <div class="flex flex-col gap-4 mt-5">
    <el-empty v-if="filteredEvents.length == 0" />
    <EventCard
      v-else
      v-for="event in paginatedEvents"
      :key="event.id"
      :event="event"
    />
  </div>

  <el-pagination
    v-model:current-page="currentPage"
    :page-size="pageSize"
    :total="totalEvents"
    layout="prev, pager, next"
    class="mt-5"
  />
</template>

<script setup lang="ts">
import { computed, onBeforeMount, ref, watch } from "vue";
import { Search } from "@element-plus/icons-vue";
import EventCard from "@/views/components/EventCard.vue";
import { getEvents } from "@/services/EventGatherService";
import { ElNotification } from "element-plus";
import { AppEventTypesEnum } from "@/utils/Constants";
import { useEventStore } from "@/stores/EventStore";
import { getUser } from "@/services/userService";

const eventStore = useEventStore();
const isLoading = ref(false);

const typeFilter = ref("");
const dateRangeFilter = ref("");
const nameSearchFilter = ref("");

const currentPage = ref(1);
const pageSize = 3;

const filteredEvents = computed(() => {
  let filtered = eventStore.appEvents;

  if (typeFilter.value != "" && typeFilter.value != null) {
    filtered = eventStore.appEvents.filter(
      (x) => x.type.name == typeFilter.value,
    );
  }

  if (dateRangeFilter.value != "" && dateRangeFilter.value != null) {
    filtered = eventStore.appEvents.filter(
      (x) =>
        x.date_start >= new Date(dateRangeFilter.value[0]) &&
        x.date_end <= new Date(dateRangeFilter.value[1]),
    );
  }

  if (nameSearchFilter.value != "") {
    filtered = eventStore.appEvents.filter((x) =>
      x.full_name.toLowerCase().includes(nameSearchFilter.value.toLowerCase()),
    );
  }

  return filtered;
});

const totalEvents = computed(() => filteredEvents.value.length);

const paginatedEvents = computed(() => {
  const start = (currentPage.value - 1) * pageSize;
  const end = start + pageSize;
  return filteredEvents.value.slice(start, end);
});

onBeforeMount(async () => {
  const data = await getUser();

  console.log({data});

  try {
    isLoading.value = true;
    eventStore.storeEvents(await getEvents());
    isLoading.value = false;
  } catch (error) {
    ElNotification({
      title: "Внимание!",
      message: "Ошибка получения событий",
      type: "warning",
    });
    isLoading.value = false;
  }
});
</script>

<style scoped></style>
