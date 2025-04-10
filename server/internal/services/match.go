package services

import (
	"fmt"
	"nashor/internal/helpers"
	"nashor/internal/types"
	"sync"
)

const matchBase = "/lol/match/v5/matches"

type MatchService struct {
	riotClient *RiotClient
}

func NewMatchService(rc *RiotClient) MatchService {
	return MatchService{
		riotClient: rc,
	}
}

func (s MatchService) GetMatchIdsByPuuid(server, puuid, start, count string) ([]string, error) {
	queries := make(map[string]string)
	var ids []string

	queries["count"] = count
	queries["start"] = start

	resp, err := s.riotClient.Get(helpers.GetRegionFromServer(server), fmt.Sprintf(matchBase+"/by-puuid/%s/ids", puuid), queries)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ids, err = helpers.ParseBody[[]string](resp.Body)

	if err != nil {
		return nil, err
	}

	return ids, nil
}

func (s MatchService) GetMatchDataById(server, id string, ch chan types.MatchDto, wg *sync.WaitGroup) {
	defer wg.Done()
	var d types.MatchDto

	m, err := s.riotClient.db.GetMatchById(id)

	if err == nil {
		ch <- m
		return
	}

	fmt.Println(m, err)

	resp, err := s.riotClient.Get(helpers.GetRegionFromServer(server), fmt.Sprintf(matchBase+"/%s", id), nil)

	if err != nil {
		return
	}

	defer resp.Body.Close()

	d, err = helpers.ParseBody[types.MatchDto](resp.Body)

	if err != nil {
		return
	}

	s.riotClient.db.CreateMatch(d)

	ch <- d
}

func (s MatchService) GetMatchDataByIds(server string, ids []string) ([]types.MatchDto, error) {
	ch := make(chan types.MatchDto)
	var wg sync.WaitGroup

	for _, id := range ids {
		wg.Add(1)
		go s.GetMatchDataById(server, id, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	matches := make([]types.MatchDto, 0)
	for r := range ch {
		matches = append(matches, r)
	}

	return matches, nil
}
