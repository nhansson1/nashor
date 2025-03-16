<script setup lang="ts">
import type { IMatch } from "@/types/match";
import { ref } from "vue";
import {
  calculateCreepscorePerMinute,
  calculateKd,
  getGoldEarnedString,
} from "@/utils/participant-utils";

import MatchDetails from "@/components/MatchDetails.vue";
import SummonerData from "@/components/SummonerData.vue";
import PlayerStats from "@/components/PlayerStats.vue";
import ItemContainer from "@/components/ItemContainer.vue";
import ParticipantContainer from "@/components/ParticipantContainer.vue";
import {
  getGameDurationString,
  getTimeElapsedSinceMatch,
  getTimeSinceMatchString,
} from "@/utils/time";

const props = defineProps<{ match: IMatch; puuid: string }>();
const isOpen = ref();
const win = props.match.info.participants[0].win;
const matchOutcome = win ? "victory" : "defeat";
const targetSummoner = props.match.info.participants.find(
  (participant) => participant.puuid === props.puuid
)!;

const targetStats = [
  targetSummoner.kills,
  targetSummoner.deaths,
  targetSummoner.assists,
].join("/");

const targetCreepScore =
  targetSummoner.neutralMinionsKilled + targetSummoner.totalMinionsKilled;

const targetCspm = calculateCreepscorePerMinute(
  targetSummoner,
  props.match.info.gameDuration
);

const targetItems = [
  targetSummoner.item0,
  targetSummoner.item1,
  targetSummoner.item2,
  targetSummoner.item6,
  targetSummoner.item3,
  targetSummoner.item4,
  targetSummoner.item5,
];

const targetCreepScoreString = `${targetCreepScore} (${targetCspm})`;

const targetVisionScoreString = `${targetSummoner.visionScore} vision`;
const targetGoldEarned = getGoldEarnedString(targetSummoner);
const targetKd = calculateKd(targetSummoner);

const timeSinceMatch = getTimeSinceMatchString(
  getTimeElapsedSinceMatch(
    props.match.info.gameStartTimestamp,
    props.match.info.gameDuration
  ),
  props.match.info.gameStartTimestamp + props.match.info.gameDuration
);

const matchDuration = getGameDurationString(props.match.info.gameDuration);

const handleOpen = () => {
  isOpen.value = !isOpen.value;
};
</script>

<template>
  <div class="match">
    <div :class="['match__bar', `match__bar--${matchOutcome}`]"></div>
    <div class="match__container">
      <MatchDetails
        :match-outcome="matchOutcome"
        :match-duration="matchDuration"
        :time-since-match="timeSinceMatch"
        class="match__details--mobile"
      />
      <div class="match__body">
        <MatchDetails
          :match-outcome="matchOutcome"
          :match-duration="matchDuration"
          :time-since-match="timeSinceMatch"
        />
        <SummonerData :champion-name="targetSummoner.championName" />
        <PlayerStats
          :stats="targetStats"
          :kd="targetKd"
          :creep-score="targetCreepScoreString"
          :vision-score="targetVisionScoreString"
          :gold-earned="targetGoldEarned"
        />
        <ItemContainer :items="targetItems" />
        <ParticipantContainer
          :participants="match.info.participants"
          :target-summoner-puuid="targetSummoner.puuid"
        />
      </div>
      <div v-if="isOpen">1</div>
    </div>
    <div @click="handleOpen" class="match__collapser"></div>
  </div>
</template>

<style scoped>
.match {
  background-color: var(--background);
  display: flex;
  height: fit-content;
  border-radius: var(--rounding-base);
  color: white;
}

.match__container {
  flex: 1;
  margin: 0.5rem;
}

.match__body {
  display: flex;
  align-items: center;
  justify-content: space-between;
  color: #ffffff67;
  font-size: 0.75rem;
}

.match__bar {
  width: 0.25rem;
  border-top-left-radius: var(--rounding-base);
  border-bottom-left-radius: var(--rounding-base);
}

.match__bar--victory {
  background-color: var(--col-accent);
}

.match__bar--defeat {
  background-color: var(--col-defeat);
}

.match__collapser {
  cursor: pointer;
  padding: 0 var(--margin-base);
}
</style>
