package types

import (
	"encoding/json"
	"fmt"
)

type ParticipantDto struct {
	ChampionName         string   `json:"championName"`
	ChampionId           int      `json:"championId"`
	Puuid                string   `json:"puuid"`
	Kills                int      `json:"kills"`
	Deaths               int      `json:"deaths"`
	Assists              int      `json:"assists"`
	NeutralMinionsKilled int      `json:"neutralMinionsKilled"`
	TotalMinionsKilled   int      `json:"totalMinionsKilled"`
	GoldEarned           int      `json:"goldEarned"`
	VisionScore          int      `json:"visionScore"`
	RiotIdGameName       string   `json:"riotIdGameName"`
	RiotIdTagline        string   `json:"riotIdTagline"`
	SummonerName         string   `json:"summonerName"`
	Win                  bool     `json:"win"`
	Item0                int      `json:"item0"`
	Item1                int      `json:"item1"`
	Item2                int      `json:"item2"`
	Item3                int      `json:"item3"`
	Item4                int      `json:"item4"`
	Item5                int      `json:"item5"`
	Item6                int      `json:"item6"`
	Summoner1Id          int      `json:"summoner1Id"`
	Summoner2Id          int      `json:"summoner2Id"`
	Perks                PerksDto `json:"perks"`
}

type PerksDto struct {
	Styles []PerkStyleDto `json:"styles"`
}

func (p *PerksDto) Scan(value interface{}) error {
	if b, ok := value.([]byte); ok {
		return json.Unmarshal(b, p)
	}

	return fmt.Errorf("unable to scan into PerksDto: %v", value)
}

type PerkStyleDto struct {
	Style      int                     `json:"style"`
	Selections []PerkStyleSelectionDto `json:"selections"`
}

type PerkStyleSelectionDto struct {
	Perk int `json:"perk"`
}

type MetadataDto struct {
	MatchId string `json:"matchId"`
}

type InfoDto struct {
	GameDuration       int              `json:"gameDuration"`
	GameEndTimestamp   int              `json:"gameEndTimestamp"`
	GameStartTimestamp int              `json:"gameStartTimestamp"`
	QueueId            int              `json:"queueId"`
	Participants       []ParticipantDto `json:"participants"`
}

type MatchDto struct {
	Metadata MetadataDto `json:"metadata"`
	Info     InfoDto     `json:"info"`
}
