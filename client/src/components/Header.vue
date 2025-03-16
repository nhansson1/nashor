<script setup lang="ts">
import SearchBar from "@/components/SearchBar.vue";
</script>

<template>
  <header class="header">
    <h1 class="header__title">nashor</h1>
    <SearchBar class="searchbar--header" />
  </header>
</template>

<style scoped>
.header {
  display: flex;
  justify-content: center;
  align-items: center;
  background-color: var(--foreground);
}

.header__title {
  color: white;
  display: none;
}
</style>
