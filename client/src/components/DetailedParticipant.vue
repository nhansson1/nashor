<script setup lang="ts">
import type { IParticipant } from "@/types/match";
import ItemContainer from "./ItemContainer.vue";
import {
    getGoldEarnedString,
    getParticipantItems,
} from "@/utils/participant-utils";

const props = defineProps<{
    participant: IParticipant;
    matchDuration: number;
}>();
const stats = [
    props.participant.kills,
    props.participant.deaths,
    props.participant.assists,
].join("/");
const items = getParticipantItems(props.participant);
const champIconSrc = `https://cdn.nashor.gg/assets/15.7.1/img/champion/${props.participant.championName}.png`;
const creepScore =
    props.participant.totalMinionsKilled +
    props.participant.neutralMinionsKilled;
</script>

<template>
    <div class="detailed-participant">
        <div class="detailed-participant__summoner">
            <div class="detailed-participant-icon-container">
                <img class="detailed-participant-icon" :src="champIconSrc" alt="champion-icon" />
            </div>
            <p class="detailed-participant__name">
                {{ participant.riotIdGameName }}
            </p>
        </div>
        <div class="detailed-participant__stats">
            <p>{{ stats }}</p>
            <p>{{ creepScore }}</p>
            <p>{{ getGoldEarnedString(participant) }}</p>
        </div>
        <ItemContainer :small="true" :foreground="false" :items="items" />
    </div>
</template>

<style scoped>
.detailed-participant {
    display: flex;
    align-items: center;
    background-color: var(--background);
    border-radius: var(--rounding-base);
    color: white;
    font-size: clamp(0.65rem, 2vw, 0.85rem);
}

.detailed-participant__summoner {
    display: flex;
    align-items: center;
}

.detailed-participant__name {
    max-height: 2ch;
    width: clamp(8ch, 10vw, 10rem);
    white-space: no-wrap;
    text-overflow: ellipsis;
    overflow: hidden;
}

.detailed-participant__stats {
    display: grid;
    margin: var(--margin-base);
    grid-template-columns: repeat(3, 1fr);
    place-items: center;
    flex: 1;
}

.detailed-participant-icon-container {
    width: 2rem;
    height: 2rem;
    margin: var(--margin-base);
    border-radius: 50%;
    overflow: hidden;
}

.detailed-participant-icon {
    height: 100%;
    object-fit: cover;
    transform: scale(1.1);
}
</style>
