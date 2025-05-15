<script setup lang="ts">
defineProps<{ iconSrc?: string }>();
</script>

<!-- modifier classes: small, medium, foreground -->
<template>
  <div class="icon">
    <img
      @mouseenter="$emit('mEnter')"
      @mouseleave="$emit('mLeave')"
      v-if="iconSrc"
      class="icon__img"
      :src="iconSrc"
      alt="icon"
    />
  </div>
</template>

<style scoped>
.icon {
  background-color: var(--foreground);
  width: clamp(1.125rem, 1vw, 1.5rem);
  height: clamp(1.125rem, 1vw, 1.5rem);
  border-radius: var(--rounding-base);
}

.icon--foreground {
  background-color: var(--background);
}

.icon--medium {
  width: clamp(1rem, 1vw, 1.225rem);
  height: clamp(1rem, 1vw, 1.225rem);
}

.icon--small {
  width: 0.8rem;
  height: 0.8rem;
}

.icon__img {
  height: 100%;
  object-fit: cover;
  border-radius: var(--rounding-base);
}
</style>
