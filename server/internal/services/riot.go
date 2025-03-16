package services

import (
	"fmt"
	"nashor/internal/helpers"
	"nashor/internal/problem"
	ratelimiter "nashor/internal/rate-limiter"
	"nashor/internal/storage"
	"net/http"
	"strings"
	"time"
)

type RiotClient struct {
    apiKey string
    region string
    limits map[string]*ratelimiter.RateLimiter

    httpClient *http.Client
    cache *storage.RedisClient
}

func NewRiotClient(apiKey, region string) *RiotClient {
    client := &http.Client{
        Timeout: 10 * time.Second,
    }

    rdb := storage.NewRedisClient()

    limits := make(map[string]*ratelimiter.RateLimiter)
    routingValues := []string{"EUROPE", "AMERICAS", "ASIA", "EUW1", "NA1", "EUN1"}

    for _, r := range routingValues {
        limits[strings.ToLower(r)] = ratelimiter.NewRateLimiter()
    }

    return &RiotClient{
        apiKey: apiKey,
        region: region,
        httpClient: client,
        limits: limits,

        cache: rdb,
    }
}

func (rc RiotClient) Get(region, endpoint string, queries map[string]string) (*http.Response, error) {
    r := strings.ToLower(region)

    if rc.limits[r].IsLimited() {
        if rc.limits[r].SlowBucket.IsEmpty() {
            return nil, problem.NewRequestError(429, "No tokens remaining retry later")
        }

        <-time.After(time.Second * 1)
        return rc.Get(region, endpoint, queries)
    }

    rc.limits[r].ConsumeToken()

    headers := make(map[string]string)
	headers["X-Riot-Token"] = rc.apiKey

    u := helpers.CreateRiotUrl(region, endpoint, queries)

	resp, err := helpers.MakeRequest(rc.httpClient, u, headers)

	if err != nil {
		err = fmt.Errorf("Request to %s failed: %w", u.String(), err)
		return nil, err
	}

	fmt.Printf("[LOG] Request to: %s returned status: %d\n", u.Path, resp.StatusCode)
	if resp.StatusCode != http.StatusOK {  
        return nil, problem.NewRequestError(resp.StatusCode, u.Path)
	}

    return resp, nil
}
