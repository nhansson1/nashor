<script setup lang="ts">
import Container from "@/components/ui/Container.vue";

const props = defineProps<{ riotId: string; profileIconId: number }>();
const iconSrc = `${import.meta.env.VITE_ASSETS_BASE}/img/profileicon/${props.profileIconId}.png`;
</script>

<template>
  <Container class="summoner-container">
    <img
      class="summoner-container__profile-icon"
      :src="iconSrc"
      alt="profile-icon"
    />
    <div class="summoner-container__account">
      <p>{{ props.riotId }}</p>
    </div>
  </Container>
</template>

<style scoped>
.summoner-container {
  display: flex;
  color: white;
}

.summoner-container__account {
  margin: var(--margin-base);
  font-size: clamp(0.85rem, 1vw, 1rem);
  min-width: 16ch;
}

.summoner-container__profile-icon {
  height: clamp(4rem, 4vw, 5rem);
  width: clamp(4rem, 4vw, 5rem);
  object-fit: contain;
}
</style>
