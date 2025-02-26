package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"nashor/internal/problem"
	"net/http"
	"net/url"
	"os"
	"time"
)

var httpClient = &http.Client{Timeout: 10 * time.Second}

func HttpResFromErr(err error) problem.ErrorResponse {
	var perr problem.ErrorResponse
	if errors.As(err, &perr) {
		return perr
	}

	return problem.InternalServerError()
}

func CreateRequest(h, p string, headers map[string]string) (*http.Response, error) {
	u := &url.URL{
		Scheme: "https",
		Host:   h,
		Path:   p,
	}

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
