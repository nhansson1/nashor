<script setup lang="ts">
const props = defineProps<{ iconSrc?: string }>();
</script>

<template>
  <div class="icon">
    <img v-if="iconSrc" class="icon__img" :src="iconSrc" alt="icon" />
  </div>
</template>

<style scoped>
.icon {
  background-color: var(--foreground);
  width: clamp(1.125rem, 1vw, 1.5rem);
  height: clamp(1.125rem, 1vw, 1.5rem);
  border-radius: var(--rounding-base);
}

.icon--medium {
  width: 0.8rem;
  height: 0.8rem;
}

.icon--small {
  width: 0.5rem;
  height: 0.5rem;
}

.icon__img {
  height: 100%;
  object-fit: cover;
  border-radius: var(--rounding-base);
}
</style>
