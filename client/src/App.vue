<script setup lang="ts">
import { RouterView } from "vue-router";
import Header from "@/components/Header.vue";
import { useRoute } from "vue-router";

const route = useRoute();
</script>

<template>
  <Header v-if="route.name != 'home'" />
  <RouterView />
</template>
