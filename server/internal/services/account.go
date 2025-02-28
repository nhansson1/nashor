package services

import (
	"fmt"
	"nashor/internal/helpers"
)

const accountBase = "/riot/account/v1/accounts"

type AccountDto struct {
	Puuid    string `json:"puuid"`
	GameName string `json:"gameName"`
	TagLine  string `json:"tagLine"`
}

func GetAccountByRiotId(region, gameName, tagLine string) (AccountDto, error) {
	u := helpers.CreateRiotUrl(region, fmt.Sprintf(accountBase+"/by-riot-id/%s/%s", gameName, tagLine), nil)

	data, err := GetEndpointJson[AccountDto](u)

	if err != nil {
		return data, err
	}

	return data, nil
}
