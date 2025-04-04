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
import { getParticipantItems } from "@/utils/participant-utils";
import {
    getGameDurationString,
    getTimeElapsedSinceMatch,
    getTimeSinceMatchString,
} from "@/utils/time";
import DetailedParticipantsList from "@/components/DetailedParticipantsList.vue";

const props = defineProps<{ match: IMatch; puuid: string }>();
const isOpen = ref();
const targetSummoner = props.match.info.participants.find(
    (participant) => participant.puuid === props.puuid
)!;
const win = targetSummoner.win;
const matchOutcome = win ? "victory" : "defeat";

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

const targetItems = getParticipantItems(targetSummoner);
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
    <div class="match-history__match-block">
        <div class="match">
            <div :class="['match__bar', `match__bar--${matchOutcome}`]"></div>
            <div class="match__container">
                <MatchDetails :match-outcome="matchOutcome" :match-duration="matchDuration"
                    :time-since-match="timeSinceMatch" :queue-id="match.info.queueId" class="match__details--mobile" />
                <div class="match__body">
                    <MatchDetails :match-outcome="matchOutcome" :match-duration="matchDuration"
                        :time-since-match="timeSinceMatch" :queue-id="match.info.queueId" />
                    <SummonerData :champion-id="targetSummoner.championId" :summoner1-id="targetSummoner.summoner1Id"
                        :summoner2-id="targetSummoner.summoner2Id"
                        :perk1-id="targetSummoner.perks.styles[0].selections[0].perk"
                        :perk2-id="targetSummoner.perks.styles[1].style" />
                    <PlayerStats :stats="targetStats" :kd="targetKd" :creep-score="targetCreepScoreString"
                        :vision-score="targetVisionScoreString" :gold-earned="targetGoldEarned" />
                    <ItemContainer :foreground="false" :items="targetItems" />
                    <ParticipantContainer :participants="match.info.participants"
                        :target-summoner-puuid="targetSummoner.puuid" />
                </div>
            </div>
            <div @click="handleOpen" class="match__collapser"></div>
        </div>
        <DetailedParticipantsList v-if="isOpen" :participants="match.info.participants"
            :match-duration="match.info.gameDuration" />
    </div>
</template>

<style scoped>
.match-history__match-block {
    background-color: var(--background);
}

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
    max-height: 100%;
}
</style>
