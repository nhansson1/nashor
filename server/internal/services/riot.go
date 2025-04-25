package services

import (
	"fmt"
	"nashor/internal/problem"
	"nashor/internal/ratelimiter"
	"nashor/internal/storage"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type RiotClient struct {
	apiKey string
	region string
	limits map[string]*ratelimiter.RateLimiter

	httpClient *http.Client
	cache      *storage.RedisClient
	db         *storage.PostgresClient
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

	pdb := storage.NewPostgresClient()

	return &RiotClient{
		apiKey:     apiKey,
		region:     region,
		httpClient: client,
		limits:     limits,

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
