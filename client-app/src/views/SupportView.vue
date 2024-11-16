<template>
  <el-tabs v-model="supportType">
    <el-tab-pane label="Государственные меры поддержки" name="government">
      <SupportList
        :items="governmentSupportExternalLinks"
        :initialPage="currentGovernmentPage"
        :pageSize="3"
        @page-change="updateGovernmentPage"
      />
    </el-tab-pane>
    <el-tab-pane label="Региональные меры поддержки" name="regional">
      <SupportList
        :items="regionalSupportExternalLinks"
        :initialPage="currentRegionalPage"
        :pageSize="3"
        @page-change="updateRegionalPage"
      />
    </el-tab-pane>
  </el-tabs>
</template>

<script lang="ts" setup>
import SupportList from "@/views/components/SupportList.vue";
import {
  governmentSupportExternalLinks,
  regionalSupportExternalLinks,
} from "@/utils/ExternalLinks";
import { ref } from "vue";

const supportType = ref("government");

const currentGovernmentPage = ref(1);
const currentRegionalPage = ref(1);

const updateGovernmentPage = (page: number) => {
  currentGovernmentPage.value = page;
};

const updateRegionalPage = (page: number) => {
  currentRegionalPage.value = page;
};
</script>
