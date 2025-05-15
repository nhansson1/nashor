<script setup lang="ts">
defineProps<{ teamOutcome: string }>();
</script>

<template>
  <div class="detailed-participant-list__table">
    <p
      :class="[
        'detailed-participant-list__team-outcome',
        `detailed-participant-list__team-outcome--${teamOutcome}`,
      ]"
    >
      {{ teamOutcome }}
    </p>
    <div class="detailed-participant-list__desc">
      <p>KDA</p>
      <p>CS</p>
      <p>GOLD</p>
    </div>
    <div></div>
  </div>
</template>

<style scoped>
.detailed-participant-list__team-outcome {
  display: flex;
  align-items: center;
  font-size: 1rem;
}

.detailed-participant-list__team-outcome--victory {
  color: var(--col-accent);
}

.detailed-participant-list__team-outcome--defeat {
  color: var(--col-defeat);
}

.detailed-participant-list__table {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  align-items: center;
  color: white;
}

.detailed-participant-list__desc {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  place-items: center;
  flex: 1;
  color: #ffffffa7;
}
</style>
