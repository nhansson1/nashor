package services

import (
	"encoding/json"
	"fmt"
	"nashor/internal/helpers"
	"time"
)

const accountBase = "/riot/account/v1/accounts"

type AccountDto struct {
	Puuid    string `json:"puuid"`
	GameName string `json:"gameName"`
	TagLine  string `json:"tagLine"`
}

type AccountService struct {
    riotClient *RiotClient
}

func NewAccountService(rc *RiotClient) AccountService {
    return AccountService{
        riotClient: rc,
    }        
}

func (s *AccountService) GetAccountByRiotId(gameName, tagLine string) (AccountDto, error) {
    var out AccountDto
    key := fmt.Sprint("account:", gameName, "#", tagLine)

    err := s.riotClient.cache.GetJson(key, &out)

    if err == nil {
        return out, nil
    }

	resp, err := s.riotClient.Get(s.riotClient.region, fmt.Sprintf(accountBase + "/by-riot-id/%s/%s", gameName, tagLine), nil)

	if err != nil {
		return out, err
	}

    defer resp.Body.Close()

    out, err = helpers.ParseBody[AccountDto](resp.Body)

    if err != nil {
        return out, err
    }

    b, err := json.Marshal(out)

    if err == nil {
        s.riotClient.cache.Set(key, string(b), 60 * time.Minute)
    }

	return out, nil
}
