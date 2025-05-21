<script setup lang="ts">
import type { IParticipant } from "@/types/match";
import Icon from "./ui/Icon.vue";
import { getChampionKeyById } from "@/utils/league-utils";
import { getParticipantDisplayName } from "@/utils/participant-utils";
import { RouterLink, useRoute } from "vue-router";

defineProps<{ participant: IParticipant }>();
const ICON_BASE = `${import.meta.env.VITE_ASSETS_BASE}/img/champion`;
const route = useRoute();
const { region } = route.params;
</script>

<template>
  <RouterLink
    :to="`/summoner/${region}/${getParticipantDisplayName(participant)}/${participant.riotIdTagline}`"
    class="participant-container__participant"
  >
    <Icon
      class="icon--small"
      :icon-src="`${ICON_BASE}/${getChampionKeyById(participant.championId)}.png`"
    />
    <p class="participant-container__name">
      {{ getParticipantDisplayName(participant) }}
    </p>
  </RouterLink>
</template>

<style scoped>
.participant-container__participant {
  text-decoration: none;
  color: #ffffff67;
}

.participant-container__participant:hover {
  text-decoration: underline;
  color: white;
}

.participant-container__name {
  width: 12ch;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  margin-left: var(--margin-small);
}

.participant-container__participant--target {
  color: white;
}

.participant-container__name {
  display: none;
}

@media screen and (min-width: 1024px) {
  .participant-container__participant {
    display: flex;
    align-items: center;
  }

  .participant-container__name {
    display: block;
  }
}
</style>
