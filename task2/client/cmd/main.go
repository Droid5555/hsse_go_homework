package main

import (
	"fmt"
	"io"
	"net/http"
)

func getVersion(client *http.Client) (string, error) {
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

//func getHardOp(client *http.Client) (bool, string, error) {
//	response, err := client.Get("http://localhost:8080/hard-op")
//	if err != nil {
//		return err
//	}
//	defer response.Body.Close()
//
//	return nil
//}

func main() {
	client := &http.Client{}

	version, err := getVersion(client)
	if err != nil {
		return
	}
	fmt.Println(version)
}

/*
TODO:
	- getHardOp
	- postDecode
	- tests
	- divide code into packages
*/
