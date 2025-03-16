type QueueType = "RANKED_FLEX_SR" | "RANKED_SOLO_5x5";

export interface Account {
  gameName: string;
  tagLine: string;
}

export interface Summoner {
  accountId: string;
  profileIconId: number;
  revisionDate: number;
  id: string;
  puuid: string;
  summonerLevel: string;
}

export interface RankEntry {
  leagueId: string;
  summonerId: string;
  puuid: string;
  queueType: QueueType;
  tier: string;
  rank: string;
  leaguePoints: number;
  wins: number;
  losses: number;
  hotStreak: boolean;
  veteran: boolean;
  freshBlood: boolean;
  inactive: boolean;
}

export interface SummonerData {
  summoner: Summoner;
  account: Account;
  ranks: RankEntry[];
}
