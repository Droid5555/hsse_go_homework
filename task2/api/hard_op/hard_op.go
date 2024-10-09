package hard_op

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

const hardOpTimeout = 15 * time.Second

func GetHardOp(client *http.Client) (bool, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), hardOpTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/hard-op", nil)
	if err != nil {
		return false, "", fmt.Errorf("could not create request: %w", err)
	}

	response, err := client.Do(req)
	if err != nil {
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			return false, "", fmt.Errorf("timed out waiting for hard_op response: %w", err)
		}
		return false, "", fmt.Errorf("failed to get from hard_op: %w", err)
	}
	defer response.Body.Close()

	r, err := io.ReadAll(response.Body)
	if err != nil {
		return false, "", fmt.Errorf("error reading response body: %v", err)
	}

	return true, string(r), nil
}
