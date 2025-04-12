<script setup lang="ts">
import type { IParticipant } from "@/types/match";
import DetailedParticipantTeam from "./DetailedParticipantTeam.vue";

const props = defineProps<{
    participants: IParticipant[];
    matchDuration: number;
}>();
const teams = [
    props.participants.slice(0, props.participants.length / 2),
    props.participants.slice(props.participants.length / 2),
];
</script>

<template>
    <div class="detailed-participant-list">
        <DetailedParticipantTeam v-for="team in teams" :participants="team" :win="team[0].win"
            :match-duration="matchDuration" />
    </div>
</template>

<style scoped>
.detailed-participant-list {
    display: flex;
    flex-direction: column;
    margin: var(--margin-base);
    padding: var(--margin-base);
    background-color: var(--foreground);
    gap: var(--margin-base);
}
</style>
