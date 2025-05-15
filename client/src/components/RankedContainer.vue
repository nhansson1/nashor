<script setup lang="ts">
import type { RankEntry } from "@/types/league";
import Container from "./ui/Container.vue";
import { calculateWinRate } from "@/utils/math";

const props = defineProps<{ rankEntry: RankEntry }>();
const gamesPlayedString = `${props.rankEntry.wins}W - ${
  props.rankEntry.losses
}L (${calculateWinRate(props.rankEntry.wins, props.rankEntry.losses)}%)`;
const rankName = `${props.rankEntry.tier[0]}${props.rankEntry.tier
  .slice(1)
  .toLowerCase()}`;

const iconSrc = `${import.meta.env.VITE_ASSETS_BASE}/rank-emblems/Rank=${rankName}.png`;
</script>

<template>
  <Container class="rank-container">
    <img class="rank-container__icon" :src="iconSrc" alt="rank-emblem" />
    <div class="rank-container__details">
      <p class="rank-container__queue-type">
        {{
          rankEntry?.queueType === "RANKED_SOLO_5x5"
            ? "Ranked Solo"
            : "Ranked Flex"
        }}
      </p>
      <p class="rank-container__league">
        {{ props.rankEntry ? props.rankEntry.tier.toLowerCase() : "Unranked" }}
        {{ props.rankEntry?.rank }}
        {{ props.rankEntry ? props.rankEntry.leaguePoints + " LP" : "" }}
      </p>
      <p class="rank-container__summary">
        {{
          props.rankEntry.wins + props.rankEntry.losses > 0
            ? gamesPlayedString
            : ""
        }}
      </p>
    </div>
  </Container>
</template>

<style scoped>
.rank-container {
  display: flex;
  align-items: center;
  color: white;
}

.rank-container__icon {
  height: clamp(4rem, 4vw, 5rem);
  width: clamp(4rem, 4vw, 5rem);
  object-fit: contain;
}

.rank-container__league,
.rank-container__summary {
  color: #ffffff87;
  font-size: clamp(0.65rem, 1vw, 0.75rem);
}

.rank-container__queue-type {
  font-size: clamp(0.85rem, 1vw, 1rem);
}

.rank-container__league {
  text-transform: capitalize;
}

.rank-container__details {
  margin: var(--margin-base);
}
</style>
