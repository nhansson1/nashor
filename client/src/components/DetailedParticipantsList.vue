<script setup lang="ts">
import type { IParticipant } from "@/types/match";
import DetailedParticipant from "@/components/DetailedParticipant.vue";

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
        <div class="detailed-participant-list__team" v-for="team in teams">
            <div>
                <p :class="[
                    'detailed-participant-list__team-outcome',
                    `detailed-participant-list__team-outcome--${team[0].win ? 'victory' : 'defeat'}`,
                ]">
                    {{ team[0].win ? "victory" : "defeat" }}
                </p>
            </div>
            <DetailedParticipant v-for="participant in team" :match-duration="props.matchDuration"
                :participant="participant" />
        </div>
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

.detailed-participant-list__team {
    display: flex;
    flex: 1;
    flex-direction: column;
    gap: var(--margin-small);
    font-size: 0.65rem;
}

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
</style>
