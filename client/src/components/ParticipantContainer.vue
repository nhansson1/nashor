<script setup lang="ts">
import type { IParticipant } from "@/types/match";
import ListParticipant from "./ListParticipant.vue";

const props = defineProps<{
  participants: IParticipant[];
  targetSummonerPuuid: string;
}>();

const teams = [
  props.participants.slice(0, props.participants.length / 2),
  props.participants.slice(props.participants.length / 2),
];
</script>

<template>
  <div class="participant-container">
    <div v-for="team in teams" class="participant-container__team">
      <ListParticipant
        v-for="participant in team"
        :participant="participant"
        :class="{
          'participant-container__participant--target':
            participant.puuid === targetSummonerPuuid,
        }"
      />
    </div>
  </div>
</template>

<style scoped>
.participant-container {
  display: flex;
  justify-content: space-between;
  gap: var(--margin-small);
}

.participant-container__team {
  display: flex;
  flex-direction: column;
  gap: var(--margin-small);
}

@media screen and (min-width: 1024px) {
  .participant-container {
    width: 12rem;
  }

  .participant-container__participant {
    display: flex;
    align-items: center;
  }

  .participant-container__name {
    display: block;
  }
}
</style>
