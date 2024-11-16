<template>
  <div class="flex flex-col gap-8">
    <el-card v-for="(option, index) in paginatedLinks" :key="index">
      <header class="font-bold text-xl">{{ option.title }}</header>
      <section class="whitespace-pre-wrap">{{ option.description }}</section>
      <footer class="flex justify-start mt-5">
        <el-button type="primary" @click="openLink(option.link)">
          подробнее
        </el-button>
      </footer>
    </el-card>

    <el-pagination
      v-model:current-page="currentPage"
      :page-size="pageSize"
      :total="totalItems"
      layout="prev, pager, next"
      class="mt-5"
    />
  </div>
</template>

<script lang="ts" setup>
import { ref, computed, watch } from "vue";

interface SupportLink {
  title: string;
  description: string;
  link: string;
}

const props = defineProps<{
  items: SupportLink[];
  initialPage?: number;
  pageSize?: number;
}>();

const emit = defineEmits(["page-change"]);

const currentPage = ref(props.initialPage || 1);
const pageSize = props.pageSize || 3;

const paginatedLinks = computed(() => {
  const start = (currentPage.value - 1) * pageSize;
  const end = start + pageSize;
  return props.items.slice(start, end);
});

const totalItems = computed(() => props.items.length);

const openLink = (link: string) => {
  window.open(link);
};

watch(currentPage, (newPage) => {
  emit("page-change", newPage);
});
</script>
