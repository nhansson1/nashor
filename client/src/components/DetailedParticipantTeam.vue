<script setup lang="ts">
import DetailedParticipant from "./DetailedParticipant.vue";
import type { IParticipant } from "@/types/match";
import DetailedParticipantDataTable from "@/components/DetailedParticipantDataTable.vue";

const props = defineProps<{
  win: boolean;
  participants: IParticipant[];
  matchDuration: number;
}>();
const teamOutcome = props.win ? "victory" : "defeat";
</script>

<template>
  <div class="detailed-participant-list__team">
    <DetailedParticipantDataTable :team-outcome="teamOutcome" />
    <DetailedParticipant
      v-for="participant in participants"
      :match-duration="matchDuration"
      :participant="participant"
    />
  </div>
</template>

<style scoped>
.detailed-participant-list__team {
  display: flex;
  flex: 1;
  flex-direction: column;
  gap: var(--margin-small);
  font-size: 0.65rem;
}
</style>
