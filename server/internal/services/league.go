package services

import (
	"fmt"
	"nashor/internal/helpers"
)

type MiniSeriesDto struct {
	Losses   int    `json:"losses"`
	Progress string `json:"progress"`
	Target   int    `json:"target"`
	Wins     int    `json:"wins"`
}

type LeagueEntryDto struct {
	LeagueId     string         `json:"leagueId"`
	SummonerId   string         `json:"summonerId"`
	Puuid        string         `json:"puuid"`
	QueueType    string         `json:"queueType"`
	Tier         string         `json:"tier"`
	Rank         string         `json:"rank"`
	LeaguePoints int            `json:"leaguePoints"`
	Wins         int            `json:"wins"`
	Losses       int            `json:"losses"`
	HotStreak    bool           `json:"hotStreak"`
	Veteran      bool           `json:"veteran"`
	FreshBlood   bool           `json:"freshBlood"`
	Inactive     bool           `json:"inactive"`
	MiniSeries   *MiniSeriesDto `json:"miniSeries,omitempty"`
}

const (
	leagueBase        = "/lol/league/v4"
	entriesBySummoner = leagueBase + "/entries/by-summoner"
)

func GetRankQueusById(region, summonerId string) ([]LeagueEntryDto, error) {
	u := helpers.CreateRiotUrl(region, fmt.Sprintf(leagueBase+"/entries/by-summoner/%s", summonerId), nil)

	data, err := GetEndpointJson[[]LeagueEntryDto](u)

	if err != nil {
		return data, err
	}

	return data, nil
}
