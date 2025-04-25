package helpers

import (
	"encoding/json"
	"fmt"
	"io"
)

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
