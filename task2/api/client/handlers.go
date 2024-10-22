package client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"hsse_go_homework/task2/tools/decode_tools"
	"io"
	"net/http"
	"time"
)

type Client struct {
	*http.Client
	host string
}

func New(host string) *Client {
	return &Client{&http.Client{}, host}
}

func (client *Client) PostDecode(s string) (string, error) {
	jsonInput := decode_tools.NewInput(s)

	jsonBytes, err := json.Marshal(jsonInput)
	if err != nil {
		return "", fmt.Errorf("could not marshal json: %w", err)
	}

	response, err := client.Post(client.host+"/decode", "application/json", bytes.NewBuffer(jsonBytes))

	if err != nil {
		return "", fmt.Errorf("could not sent post request to /decode: %w", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return "", fmt.Errorf("could not read response body: %w", err)
	}

	var jsonOutput decode_tools.Output
	err = json.Unmarshal(body, &jsonOutput)
	if err != nil {
		return "", fmt.Errorf("could not unmarshal response body: %w", err)
	}

	return jsonOutput.Response, nil
}

const hardOpTimeout = 15 * time.Second

func (client *Client) GetHardOp() (bool, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), hardOpTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", client.host+"/hard-op", nil)
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

func (client *Client) GetVersion() (string, error) {
	response, err := client.Get(client.host + "/version")
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	v, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	return string(v), nil
}
