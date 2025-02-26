package services

import (
	"fmt"
	"nashor/internal/helpers"
	"nashor/internal/problem"
	"net/http"
	"os"
)

func GetEndpointJson[T any](region, endpoint string) (T, error) {
	var out T
	headers := make(map[string]string)
	headers["X-Riot-Token"] = os.Getenv("API_KEY")

	resp, err := helpers.CreateRequest(fmt.Sprint(region, ".api.riotgames.com"), endpoint, headers)

	if err != nil {
		err = fmt.Errorf("Request to %s failed: %w", endpoint, err)
		return out, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("[NASHOR-LOG] Request to: %s retured status: %d\n", endpoint, resp.StatusCode)
		if resp.StatusCode == 404 {
			return out, problem.NewErrorResponse(404, "Data not found.")
		}

		return out, problem.InternalServerError()
	}

	out, err = helpers.ParseBody[T](resp.Body)

	if err != nil {
		return out, err
	}

	return out, nil
}
