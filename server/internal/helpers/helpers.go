package helpers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

var httpClient = &http.Client{Timeout: 10 * time.Second}

func CreateRiotUrl(base, endpoint string, queries map[string]string) *url.URL {
	u := &url.URL{
		Scheme: "https",
		Host:   fmt.Sprint(base, ".api.riotgames.com"),
		Path:   endpoint,
	}

	q := u.Query()

	for k, v := range queries {
		q.Add(k, v)
	}

	u.RawQuery = q.Encode()

	return u
}

func MakeRequest(u *url.URL, headers map[string]string) (*http.Response, error) {
	req, err := http.NewRequest("GET", u.String(), nil)

	if err != nil {
		fmt.Println("Failed to create request.")
		return nil, err
	}

	req.Header.Add("X-Riot-Token", os.Getenv("API_KEY"))

	resp, err := httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func ParseBody[T any](respBody io.ReadCloser) (T, error) {
	var out T

	body, err := io.ReadAll(respBody)

	if err != nil {
		return out, fmt.Errorf("Failed to read body: %w", err)
	}

	err = json.Unmarshal(body, &out)

	if err != nil {
		return out, fmt.Errorf("Failed to parse JSON: %w", err)
	}

	return out, nil
}
