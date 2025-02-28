package services

import (
	"fmt"
	"os"
	"net/url"
	"net/http"
	"nashor/internal/helpers"
	"nashor/internal/problem"
)

func GetEndpointJson[T any](u *url.URL) (T, error) {
	var out T
	headers := make(map[string]string)
	headers["X-Riot-Token"] = os.Getenv("API_KEY")

	resp, err := helpers.MakeRequest(u, headers)

	if err != nil {
		err = fmt.Errorf("Request to %s failed: %w", u.String(), err)
		return out, err
	}

	defer resp.Body.Close()

	fmt.Printf("[NASHOR-LOG] Request to: %s returned status: %d\n", u.Path, resp.StatusCode)
	if resp.StatusCode != http.StatusOK {
		return out, problem.NewRequestError(resp.StatusCode, u.Path)
	}

	out, err = helpers.ParseBody[T](resp.Body)

	if err != nil {
		return out, err
	}

	return out, nil
}
