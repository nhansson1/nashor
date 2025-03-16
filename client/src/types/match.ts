export interface IParticipant {
  championName: string;
  puuid: string;
  kills: number;
  deaths: number;
  assists: number;
  neutralMinionsKilled: number;
  totalMinionsKilled: number;
  goldEarned: number;
  visionScore: number;
  riotIdGameName: string;
  win: boolean;
  item0: number;
  item1: number;
  item2: number;
  item3: number;
  item4: number;
  item5: number;
  item6: number;
}

interface IMetaData {
  matchId: string;
}

interface Info {
  gameDuration: number;
  gameEndTimestamp: number;
  gameStartTimestamp: number;
  participants: IParticipant[];
}

export interface IMatch {
  metadata: IMetaData;
  info: Info;
}
