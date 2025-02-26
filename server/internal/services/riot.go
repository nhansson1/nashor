package services

import (
	"fmt"
	"nashor/internal/helpers"
	"nashor/internal/problem"
	"os"
)

const riotBase = "europe.api.riotgames.com"

func GetEndpointJson[T any](endpoint string) (T, error) {
    var out T
    headers := make(map[string]string)
    headers["X-Riot-Token"] = os.Getenv("API_KEY")

    resp, err := helpers.CreateRequest(endpoint, headers)

    if err != nil {
        err = fmt.Errorf("Request to %s failed: %w", endpoint, err)
        return out, err
    }

    defer resp.Body.Close()

    if resp.StatusCode != 200 {
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
