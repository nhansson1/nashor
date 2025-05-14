package services

import (
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"log"
	"nashor/internal/problem"
	"nashor/internal/storage"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

//go:embed scripts/tokenbucket.lua
var luaStr string
var tokenBucketScript = redis.NewScript(luaStr)

type RedisBucketRes struct {
	Status string `json:"status"`
	Fast   int    `json:"fast"`
	Slow   int    `json:"slow"`
}

type RiotClient struct {
	apiKey string
	region string

	httpClient *http.Client
	cache      *storage.RedisClient
	db         *storage.PostgresClient
}

func NewRiotClient(apiKey, region string) *RiotClient {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	rdb := storage.NewRedisClient()

	pdb := storage.NewPostgresClient()

	return &RiotClient{
		apiKey:     apiKey,
		region:     region,
		httpClient: client,

		cache: rdb,
		db:    pdb,
	}
}

func (rc RiotClient) GetRegionFromServer(server string) string {
	var region string

	switch strings.ToUpper(server) {
	case "EUW1", "EUNE", "TR", "ME1", "RU":
		region = "EUROPE"
	case "NA1", "BR", "LAN", "LAS":
		region = "AMERICAS"
	case "KR", "JP":
		region = "ASIA"
	case "OCE", "SG2", "TW2", "VN2":
		region = "SEA"
	}

	return region
}

func (rc RiotClient) createRiotUrl(routingValue, endpoint string, queries map[string]string) *url.URL {
	u := &url.URL{
		Scheme: "https",
		Host:   fmt.Sprint(routingValue, ".api.riotgames.com"),
		Path:   endpoint,
	}

	q := u.Query()

	for k, v := range queries {
		q.Add(k, v)
	}

	u.RawQuery = q.Encode()

	return u
}

func (rc RiotClient) checkRateLimits(region string) RedisBucketRes {
	res, err := tokenBucketScript.Run(context.Background(), rc.cache.GetClient(), []string{fmt.Sprint("api:rate_limits_", region)}).Result()

	if err != nil {
		log.Fatal("failed to run lua script", err)
	}

	jsonStr, ok := res.(string)

	if !ok {
		log.Fatal("lua script returned unexpected type")
	}

	var out RedisBucketRes
	json.Unmarshal([]byte(jsonStr), &out)

	return out
}

func (rc RiotClient) makeRequest(u *url.URL) (*http.Response, error) {
	req, err := http.NewRequest("GET", u.String(), nil)

	if err != nil {
		fmt.Println("Failed to create request.")
		return nil, err
	}

	req.Header.Add("X-Riot-Token", rc.apiKey)

	resp, err := rc.httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (rc RiotClient) Get(region, endpoint string, queries map[string]string) (*http.Response, error) {
	limits := rc.checkRateLimits(strings.ToLower(region))

	switch limits.Status {
	case "fast_limit":
		<-time.After(time.Second * 1)
		return rc.Get(region, endpoint, queries)
	case "slow_limit":
		return nil, problem.NewRequestError(429, "No tokens remaining retry later")
	}

	headers := make(map[string]string)
	headers["X-Riot-Token"] = rc.apiKey

	u := rc.createRiotUrl(region, endpoint, queries)

	resp, err := rc.makeRequest(u)

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
