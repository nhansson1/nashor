<script setup lang="ts">
import SearchBar from "@/components/SearchBar.vue";
</script>

<template>
    <main class="home">
        <h1 class="home__header">nashor.gg</h1>
        <SearchBar />
    </main>
</template>

<style scoped>
.home {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    height: 100vh;
}

.home__header {
    font-size: 2rem;
    color: white;
    text-transform: uppercase;
}

.home__header {
    font-size: clamp(2rem, 5vw, 3rem);
}
</style>
