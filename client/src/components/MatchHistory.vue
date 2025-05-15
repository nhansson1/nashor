<script setup lang="ts">
import { ref, watch } from "vue";
import { useRoute } from "vue-router";

import Match from "@/components/Match.vue";
import Button from "@/components/ui/Button.vue";
import type { IMatch } from "@/types/match";

const route = useRoute();
const matches = ref<IMatch[]>([]);
const props = defineProps<{ puuid: string }>();

const getMatches = async (start: number) => {
  const count = 10;
  const { region } = route.params;

  const url = new URL(
    `${import.meta.env.VITE_API_BASE}/matches/by-puuid/${region}/${props.puuid}/${start}/${count}`,
  );

  try {
    const res = await fetch(url);

    if (!res.ok) {
      return;
    }

    const data: IMatch[] = await res.json();

    matches.value = [
      ...matches.value,
      ...data.sort((a, b) => b.info.gameEndTimestamp - a.info.gameEndTimestamp),
    ];
  } catch (err) {
    console.log(err);
  }
};

watch(
  () => props.puuid,
  () => {
    matches.value = [];
    getMatches(0);
  },
  {
    immediate: true,
  },
);
</script>

<template>
  <div class="match-history">
    <Match
      v-if="matches.length > 0"
      v-for="match in matches"
      :match="match"
      :puuid="puuid"
      :key="match.metadata.matchId"
    />
    <div class="match-history__match--loading" v-else v-for="n in 10"></div>
    <Button @button-click="() => getMatches(matches.length)" text="load more" />
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

.match-history .button:last-child {
  align-self: center;
}

.match-history__match--loading {
  height: 6.375rem;
  background-color: var(--background);
  border-radius: var(--rounding-base);
  animation: skeleton ease 1.5s infinite;
}

@keyframes skeleton {
  0% {
    opacity: 1;
  }

  50% {
    opacity: 0.5;
  }

  100% {
    opacity: 1;
  }
}
</style>
