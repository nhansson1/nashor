<script setup lang="ts">
import { ref, watch } from "vue";
import { useRoute } from "vue-router";

import Match from "@/components/Match.vue";
import type { IMatch } from "@/types/match";

const route = useRoute();
const matches = ref<IMatch[]>([]);
const props = defineProps<{ puuid: string }>();

const getMatches = async (start: number) => {
    const count = 10;
    const { region } = route.params;

    const url = new URL(
        `http://localhost:8080/api/v1/matches/by-puuid/${region}/${props.puuid}/${start}/${count}`
    );

    try {
        const res = await fetch(url);

        if (!res.ok) {
            return;
        }

        const data: IMatch[] = await res.json();

        matches.value = [
            ...matches.value,
            ...data.sort(
                (a, b) => b.info.gameEndTimestamp - a.info.gameEndTimestamp
            ),
        ];
    } catch (err) {
        console.log(err);
    }
};

watch(() => props.puuid, getMatches.bind(null, 0), {
    immediate: true,
});
</script>

<template>
    <div class="match-history">
        <Match
            v-for="match in matches"
            :match="match"
            :puuid="puuid"
            :key="match.metadata.matchId"
        />
    </div>
</template>

<style scoped>
.match-history {
    display: flex;
    flex-direction: column;
    background-color: var(--foreground);
    margin-top: 0.5rem;
    flex: 1;
    padding: 0.5rem;
    gap: 0.5rem;
}
</style>
