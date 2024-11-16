<template>
  <el-tabs v-model="supportType">
    <el-tab-pane label="Федеральные меры поддержки" name="government">
      <SupportList
        :items="federalSupportExternalLinks"
        :initialPage="currentFederalPage"
        :pageSize="3"
        @page-change="updateFederalPage"
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
  federalSupportExternalLinks,
  regionalSupportExternalLinks,
} from "@/utils/ExternalLinks";
import { ref } from "vue";

const supportType = ref("government");

const currentFederalPage = ref(1);
const currentRegionalPage = ref(1);

const updateFederalPage = (page: number) => {
  currentFederalPage.value = page;
};

const updateRegionalPage = (page: number) => {
  currentRegionalPage.value = page;
};
</script>
