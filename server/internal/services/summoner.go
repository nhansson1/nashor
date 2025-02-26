package services

import (
	"fmt"
	"nashor/internal/problem"
)

type SummonerDto struct {
	AccountId     string `json:"accountId"`
	ProfileIconId int    `json:"profileIcondId"`
	RevisionDate  int    `json:"revisionDate"`
	SummonerId    string `json:"id"`
	Puuid         string `json:"puuid"`
	SummonerLevel int    `json:"summonerLevel"`
}

const (
	summonerBase    = "/lol/summoner/v4/summoners"
	summonerByPuuid = summonerBase + "/by-puuid"
)

func GetSummonerByPuuid(region, puuid string) (SummonerDto, error) {
	endpoint := fmt.Sprint(summonerByPuuid, "/"+puuid)
	data, err := GetEndpointJson[SummonerDto](region, endpoint)

	if err != nil {
		return data, problem.InternalServerError()
	}

	return data, nil
}
