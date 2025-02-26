package services

import (
	"fmt"
	"nashor/internal/problem"
)

const summonerBase = "/lol/summoner/v4/summoners"

type SummonerDto struct {
    AccountId string `json:"accountId"`
    ProfileIconId int `json:"profileIcondId"`
    RevisionDate int `json:"revisionDate"`
    SummonerId string `json:"id"`
    Puuid string `json:"puuid"`
    SummonerLevel int `json:"summonerLevel"`
}

func GetSummonerByPuuid(puuid string) (SummonerDto, error) {
    data, err := GetEndpointJson[SummonerDto](fmt.Sprint("https://euw1.api.riotgames.com", summonerBase, "/by-puuid/", puuid))

    if err != nil {    
        return data, problem.InternalServerError()
    }

    return data, nil
}
