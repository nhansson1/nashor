package services

import (
	"fmt"
	"nashor/internal/helpers"
)

type SummonerDto struct {
	AccountId     string `json:"accountId"`
	ProfileIconId int    `json:"profileIcondId"`
	RevisionDate  int    `json:"revisionDate"`
	SummonerId    string `json:"id"`
	Puuid         string `json:"puuid"`
	SummonerLevel int    `json:"summonerLevel"`
}

const summonerBase = "/lol/summoner/v4/summoners"

func GetSummonerByPuuid(region, puuid string) (SummonerDto, error) {
	u := helpers.CreateRiotUrl(region, fmt.Sprintf(summonerBase+"/by-puuid/%s", puuid), nil)

	data, err := GetEndpointJson[SummonerDto](u)

	if err != nil {
		return data, err
	}

	return data, nil
}
