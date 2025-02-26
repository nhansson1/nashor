package services

import (
	"fmt"
)

const (
	accountBase     = "/riot/account/v1/accounts"
	accountByRiotId = accountBase + "/by-riot-id"
)

type AccountDto struct {
	Puuid    string `json:"puuid"`
	GameName string `json:"gameName"`
	TagLine  string `json:"tagLine"`
}

func GetAccountByRiotId(region, gameName, tagLine string) (AccountDto, error) {
	endpoint := fmt.Sprint(accountByRiotId, "/"+gameName, "/"+tagLine)
	data, err := GetEndpointJson[AccountDto](region, endpoint)

	if err != nil {
		return data, err
	}

	return data, nil
}
