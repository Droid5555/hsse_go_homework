package version

import (
	"io"
	"net/http"
)

func GetVersion(client *http.Client) (string, error) {
	response, err := client.Get("http://localhost:8080/version")
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
