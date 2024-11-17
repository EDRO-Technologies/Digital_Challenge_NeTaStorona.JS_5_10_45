<template>
  <div class="flex flex-col">
    <div class="flex justify-between items-start h-[8vh]">
      <div class="flex items-center gap-2">
        <el-avatar />
        <p class="font-bold text-xl">IT-Югра</p>
      </div>

      <el-avatar
        v-if="!isProfilePage"
        class="flex flex-col justify-end items-end gap-2 w-48 hover:cursor-pointer"
        :icon="UserFilled"
        @click="openProfile"
      />
    </div>

    <div class="custom-style">
      <el-segmented
        v-if="!isProfilePage"
        v-model="currentPage"
        :options="navigationOptions.map((option) => option.name)"
        @change="handleNavigation"
      />
    </div>
    <div></div>
  </div>
</template>

<script setup lang="ts">
import { computed, ref } from "vue";
import { navigationOptions } from "@/utils/NavigationOptions";
import { useRoute, useRouter } from "vue-router";
import { UserFilled } from "@element-plus/icons-vue";

const router = useRouter();
const route = useRoute();
const isProfilePage = computed(() => route.path === "/profile");
const currentPage = ref(navigationOptions[0].name);

const handleNavigation = (selectedOption: string) => {
  const route = navigationOptions.find(
    (option) => option.name === selectedOption,
  )?.path;
  if (route) {
    router.push(route);
  }
};

const openProfile = () => router.push("/profile");
</script>

<style scoped>
.custom-style .el-segmented {
  --el-segmented-item-selected-color: white;
  --el-segmented-bg-color: transparent;
  --el-border-radius-base: 8px;
}
</style>
