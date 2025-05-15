<script setup lang="ts">
import { useRoute } from "vue-router";

const route = useRoute();

const { search, region } = route.query;
</script>

<template>
  <main class="not-found">
    <div class="not-found__container">
      <h1 class="not-found__title">404 - player not found</h1>
      <p class="not-found__details">
        could not find <span class="not-found__player"> {{ search }}</span> on
        region {{ region }}
      </p>
      <RouterLink class="not-found__home-link" to="/">Home</RouterLink>
    </div>
  </main>
</template>

<style scoped>
.not-found {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  background-color: var(--background);
}

.not-found__title {
  color: white;
  font-size: 1.5rem;
  text-transform: capitalize;
}

.not-found__details {
  color: #cecece;
}

.not-found__player {
  color: white;
}

.not-found__container {
  color: white;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  gap: var(--margin-base);
}

.not-found__home-link {
  color: #cecece;
}

.not-found__home-link:hover {
  color: white;
}
</style>
