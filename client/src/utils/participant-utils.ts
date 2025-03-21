import type { IParticipant } from "@/types/match";

const SECONDS_IN_MINUTE = 60;

export const getGoldEarnedString = (participant: IParticipant): string => {
  const { goldEarned } = participant;

  if (goldEarned < 1000) return goldEarned.toString();

  return `${(goldEarned / 1000).toFixed(2)}k`;
};

export const calculateKd = (participant: IParticipant): string => {
  const { kills, deaths, assists } = participant;

  if (!deaths) return "Perfect";

  return ((kills + assists) / deaths).toFixed(2);
};

export const calculateCreepscorePerMinute = (
  participant: IParticipant,
  gameLength: number
): string => {
  const creepscore =
    participant.totalMinionsKilled + participant.neutralMinionsKilled;

  const gameLengtMin = gameLength / SECONDS_IN_MINUTE;

  return (creepscore / gameLengtMin).toFixed(2);
};

export const getParticipantItems = (participant: IParticipant): number[] => {
  return [
    participant.item0,
    participant.item1,
    participant.item2,
    participant.item6,
    participant.item3,
    participant.item4,
    participant.item5,
  ];
};
