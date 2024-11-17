

<template>
  <div>
    Идет авторизация...
  </div>
</template>

<script setup lang="ts">
import { onBeforeMount } from "vue";
import router from "@/router/router";

function parseTokenFromUrl() {
    const currentUrl = window.location.href;

    const url = new URL(currentUrl);

    const params = new URLSearchParams(url.search);

    return params.get('token') as string;
  }

  onBeforeMount(() => {
    const token = parseTokenFromUrl();
    console.log({token});

    localStorage.setItem('token', token);

    router.push({path: '/'})
  })
</script>