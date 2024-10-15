package decode

import (
	"bytes"
	"encoding/json"
	"fmt"
	"hsse_go_homework/task2/tools/decode_tools"
	"io"
	"net/http"
)

func PostDecode(client *http.Client, s string) (string, error) {
	jsonInput := decode_tools.NewInput(s)

	jsonBytes, err := json.Marshal(jsonInput)
	if err != nil {
		return "", fmt.Errorf("could not marshal json: %w", err)
	}

	response, err := client.Post("http://localhost:8080/decode", "application/json", bytes.NewBuffer(jsonBytes))

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
