<script setup lang="ts">
import type { IParticipant } from "@/types/match";
import { getChampionKeyById } from "@/utils/league-utils";
import Icon from "./ui/Icon.vue";
const props = defineProps<{
  participants: IParticipant[];
  targetSummonerPuuid: string;
}>();

const team1 = props.participants.slice(0, props.participants.length / 2);
const team2 = props.participants.slice(props.participants.length / 2);

const ICON_BASE = `${import.meta.env.VITE_ASSETS_BASE}/img/champion`;
</script>

<template>
  <div class="participant-container">
    <div class="participant-container__team">
      <div
        class="participant-container__participant"
        v-for="participant in team1"
      >
        <Icon
          class="icon--small"
          :icon-src="`${ICON_BASE}/${getChampionKeyById(participant.championId)}.png`"
        />
        <p
          :class="{
            'participant-container__name': true,
            'participant-container__participant--target':
              participant.puuid === targetSummonerPuuid,
          }"
        >
          {{ participant.riotIdGameName }}
        </p>
      </div>
    </div>
    <div class="participant-container__team">
      <div
        class="participant-container__participant"
        v-for="participant in team2"
      >
        <Icon
          class="icon--small"
          :icon-src="`${ICON_BASE}/${getChampionKeyById(participant.championId)}.png`"
        />
        <p
          :class="{
            'participant-container__name': true,
            'participant-container__participant--target':
              participant.puuid === targetSummonerPuuid,
          }"
        >
          {{ participant.riotIdGameName }}
        </p>
      </div>
    </div>
  </div>
</template>

<style scoped>
.participant-container {
  display: flex;
  justify-content: space-between;
  gap: var(--margin-small);
}

.participant-container__team {
  display: flex;
  flex-direction: column;
  gap: var(--margin-small);
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
  .participant-container {
    width: 12rem;
  }

  .participant-container__participant {
    display: flex;
    align-items: center;
  }

  .participant-container__name {
    display: block;
  }
}
</style>
