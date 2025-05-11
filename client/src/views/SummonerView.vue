<script setup lang="ts">
import SummonerProfile from "@/components/SummonerProfile.vue";
import { ref, watch } from "vue";
import { useRoute, useRouter } from "vue-router";
import type { Account, Summoner, RankEntry } from "@/types/league";
import MatchHistory from "@/components/MatchHistory.vue";

interface ProfileResponse {
    account: Account;
    summoner: Summoner;
    ranks: RankEntry[];
}

const router = useRouter();
const route = useRoute();
const profile = ref<ProfileResponse>();

const fetchProfile = async () => {
    const { region, gameName, tagLine } = route.params;
    const url = new URL(
        `http://localhost:8080/api/v1/summoner/profile/by-riot-id/${region}/${gameName}/${tagLine}`
    );

    try {
        const res = await fetch(url);

        if (!res.ok) {
            router.push({
                path: "/404",
                query: { search: `${gameName}#${tagLine}`, region },
            });
            return;
        }

        const data: ProfileResponse = await res.json();

        data.ranks.sort((a) => (a.queueType == "RANKED_SOLO_5x5" ? -1 : 1));

        profile.value = data;
    } catch (err) {
        console.log(err);
    }
};

watch(() => route.params.gameName, fetchProfile, { immediate: true });
</script>

<template>
    <div class="page">
        <main v-if="profile" class="page__content">
            <SummonerProfile :summoner-data="profile" :key="profile.summoner.accountId" />
            <MatchHistory :puuid="profile.summoner.puuid" />
        </main>
    </div>
</template>

<style scoped>
.page {
    min-height: 100vh;
    background-color: var(--background-color);
    padding: 0.5rem;
}

.page__content {
    height: 100%;
}

@media screen and (min-width: 1024px) {
    .page {
        display: flex;
        justify-content: center;
    }

    .page__content {
        max-width: 60vw;
    }
}
</style>
