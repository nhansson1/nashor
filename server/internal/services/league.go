package services

import "fmt"

type MiniSeriesDto struct {
	Losses   int
	Progress string
	Target   int
	Wins     int
}

type LeagueEntryDto struct {
	LeagueId     string
	SummonerId   string
	Puuid        string
	Tier         string
	Rank         string
	LeaguePoints int
	Wins         int
	Losses       int
	HotStreak    bool
	Veteran      bool
	FreshBlood   bool
	Inactive     bool
	MiniSeries   MiniSeriesDto
}

func GetRankQueusById(summonerId string) ([]LeagueEntryDto, error) {
	data, err := GetEndpointJson[[]LeagueEntryDto](fmt.Sprintf("https://euw1.api.riotgames.com/lol/league/v4/entries/by-summoner/%s", summonerId))

	if err != nil {
		return data, err
	}

	return data, nil
}
