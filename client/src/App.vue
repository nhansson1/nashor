<script setup lang="ts">
import { RouterView } from "vue-router";
import Header from "@/components/Header.vue";
import Footer from "./components/Footer.vue";
import { useRoute } from "vue-router";

const route = useRoute();
console.log(route);
</script>

<template>
    <Header v-if="route.name && route.name != 'home'" />
    <RouterView />
    <Footer />
</template>
