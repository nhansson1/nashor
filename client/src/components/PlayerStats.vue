<script setup lang="ts">
const props = defineProps<{
  stats: string;
  kd: string;
  creepScore: string;
  visionScore: string;
  goldEarned: string;
}>();
</script>

<template>
  <div class="match__player-stats">
    <p>{{ stats }}</p>
    <p>
      {{ kd }}
      KD
    </p>
    <p>{{ creepScore }}</p>
    <p>{{ visionScore }}</p>
    <p>{{ goldEarned }}</p>
  </div>
</template>

<style scoped>
.match__player-stats {
  text-align: center;
  color: #ffffff87;
}
</style>
