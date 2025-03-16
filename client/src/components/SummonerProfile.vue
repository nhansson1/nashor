<script setup lang="ts">
import type { SummonerData } from "@/types/league";
import RankedContainer from "@/components/RankedContainer.vue";
import SummonerContainer from "@/components/SummonerContainer.vue";

const props = defineProps<{ summonerData: SummonerData }>();
const riotId = `${props.summonerData.account.gameName}#${props.summonerData.account.tagLine}`;
</script>

<template>
  <div class="summoner-profile">
    <SummonerContainer
      :profile-icon-id="props.summonerData.summoner.profileIconId"
      :riot-id="riotId"
    />
    <RankedContainer :rank-entry="props.summonerData.ranks[0]" />
    <RankedContainer :rank-entry="props.summonerData.ranks[1]" />
  </div>
</template>

<style scoped>
.summoner-profile {
  display: grid;
  align-content: center;
  justify-content: space-between;
  grid-template-columns: repeat(2, 1fr);
  gap: var(--margin-base);
}

.summoner-container {
  grid-column: span 2;
}

@media screen and (min-width: 1024px) {
  .summoner-profile {
    display: flex;
    align-items: center;
    background-color: var(--foreground);
    grid-template-columns: repeat(3, 1fr);
  }

  .summoner-container {
    grid-column: span 1;
  }
}
</style>
