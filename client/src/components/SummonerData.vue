<script setup lang="ts">
import { getChampionKeyById } from '@/utils/league-utils';
import Icon from './ui/Icon.vue';
import { getPerkSource, getSummonerSource } from '@/utils/participant-utils';

const props = defineProps<{ championId: number, summoner1Id: number, summoner2Id: number, perk1Id: number, perk2Id: number }>();
const championIcon = `https://cdn.nashor.gg/assets/15.7.1/img/champion/${getChampionKeyById(props.championId)}.png`;
const summoner1 = getSummonerSource(props.summoner1Id);
const summoner2 = getSummonerSource(props.summoner2Id);
const perk1 = getPerkSource(props.perk1Id);
const perk2 = getPerkSource(props.perk2Id);

const summonerSources = [summoner1, summoner2].map(summoner => {
    if (!summoner)
        return "";

    return `https://cdn.nashor.gg/assets/15.7.1/img/spell/${summoner}`;
});

const perkSources = [perk1, perk2].map(perk => {
    if (!perk)
        return "";

    return `https://cdn.nashor.gg/assets/img/${perk}`
});

</script>

<template>
    <div class="match__summoner-data">
        <div class="match__champ-icon-container">
            <img class="match__champ-icon" :src="championIcon" alt="champion-icon" />
        </div>
        <div class="match__summoner-icons">
            <Icon class="icon--medium" :icon-src="summonerSources[0]" />
            <Icon class="icon--foreground icon--medium" :icon-src="perkSources[0]" />
            <Icon class="icon--medium" :icon-src="summonerSources[1]" />
            <Icon class="icon--foreground icon--medium" :icon-src="perkSources[1]" />
        </div>
    </div>
</template>

<style scoped>
.match__summoner-data {
    display: flex;
    align-items: center;
}

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

.match__summoner-icons {
    display: grid;
    grid-template-columns: repeat(2, 1fr);
    gap: var(--margin-small);
}

.match__summoner-icon-container {
    width: 1.25rem;
    height: 1.25rem;
}
</style>
