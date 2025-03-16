<script setup lang="ts">
const props = defineProps<{ championName: string }>();
const championIcon = `https://cdn.nashor.gg/assets/15.5.1/img/champion/${props.championName}.png`;
</script>

<template>
  <div class="match__summoner-data">
    <div class="match__champ-icon-container">
      <img class="match__champ-icon" :src="championIcon" alt="champion-icon" />
    </div>
  </div>
</template>

<style scoped>
.match__champ-icon-container {
  width: 3rem;
  height: 3rem;
  margin: var(--margin-base);
  border-radius: 50%;
  overflow: hidden;
}

.match__champ-icon {
  height: 100%;
  object-fit: cover;
  transform: scale(1.1);
}
</style>
