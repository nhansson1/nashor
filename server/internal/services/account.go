package services

import (
	"fmt"
)

const accountBase = "/riot/account/v1/accounts"
const accountByRiotId = accountBase + "/by-riot-id"

type AccountDto struct {
    Puuid string `json:"puuid"`
    GameName string `json:"gameName"`
    TagLine string `json:"tagLine"`
}

func GetAccountByRiotId(gameName, tagLine string) (AccountDto, error) {
    data, err := GetEndpointJson[AccountDto](fmt.Sprint("https://", riotBase, accountByRiotId, "/", gameName, "/", tagLine))

    if err != nil {
        return data, err
    } 

    return data, nil
}
