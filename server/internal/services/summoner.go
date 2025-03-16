package services

import (
	"fmt"
	"nashor/internal/helpers"
)

type SummonerDto struct {
	AccountId     string `json:"accountId"`
	ProfileIconId int    `json:"profileIconId"`
	RevisionDate  int    `json:"revisionDate"`
	SummonerId    string `json:"id"`
	Puuid         string `json:"puuid"`
	SummonerLevel int    `json:"summonerLevel"`
}

const summonerBase = "/lol/summoner/v4/summoners"

type SummonerService struct {
    riotClient *RiotClient
}

func NewSummonerService(rc *RiotClient) SummonerService {
    return SummonerService{
        riotClient: rc,
    }
}

func (s SummonerService) GetSummonerByPuuid(region, puuid string) (SummonerDto, error) {
    var summoner SummonerDto

	resp, err := s.riotClient.Get(region, fmt.Sprintf(summonerBase + "/by-puuid/%s", puuid), nil)

	if err != nil {
		return summoner, err
	}

    defer resp.Body.Close()

    summoner, err = helpers.ParseBody[SummonerDto](resp.Body)

    if err != nil {
        return summoner, err 
    }

	return summoner, nil
}
