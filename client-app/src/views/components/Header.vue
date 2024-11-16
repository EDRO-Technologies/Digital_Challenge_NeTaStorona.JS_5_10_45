<template>
  <div class="flex flex-col">
    <div class="flex justify-between items-start">
      <div class="flex items-center gap-2">
        <el-avatar />
        <p class="font-bold text-xl">Цифровое пространство</p>
      </div>

      <div class="flex flex-col justify-end items-end gap-2 w-48">
        <div class="flex flex-col items-center">
          <el-avatar size="large" class="avatar" />
          <p>Тест Тестович</p>
        </div>
      </div>
    </div>

    <div class="custom-style">
      <el-segmented
        v-model="currentPage"
        :options="navigationOptions.map((option) => option.name)"
        @change="handleNavigation"
      />
    </div>
    <div></div>
  </div>
</template>

<script setup lang="ts">
import { ref } from "vue";
import { navigationOptions } from "@/utils/NavigationOptions";
import { useRouter } from "vue-router";

const router = useRouter();
const currentPage = ref(navigationOptions[0].name);

const handleNavigation = (selectedOption: string) => {
  const route = navigationOptions.find(
    (option) => option.name === selectedOption,
  )?.path;
  if (route) {
    router.push(route);
  }
};
</script>

<style scoped>
.custom-style .el-segmented {
  --el-segmented-item-selected-color: var(--el-text-color-primary);
  --el-segmented-item-selected-bg-color: #00ffc4;
  --el-segmented-bg-color: transparent;
  --el-border-radius-base: 16px;
}
</style>
