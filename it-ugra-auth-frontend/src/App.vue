<template>
  <div id="app">
    <nav class="mt-4 absolute top-0 left-4 flex flex-row bg-white px-4 py-2 rounded-md border gap-4">
      <router-link to="/">Форма фхода</router-link>
      <router-link to="/agentreqs">Панель администратора</router-link>
    </nav>
    <router-view />
  </div>
</template>

<script setup lang="ts">

</script>
