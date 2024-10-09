package version_tools

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Version struct {
	Major int `json:"VersionMajor"`
	Minor int `json:"VersionMinor"`
	Patch int `json:"VersionPatch"`
}

var VERSION Version

func LoadFromJson(filePath string) error {
	jsonFile, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("error opening JSON file: %w", err)
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return fmt.Errorf("error reading JSON file: %w", err)
	}

	err = json.Unmarshal(byteValue, &VERSION)
	if err != nil {
		return fmt.Errorf("error unmarshalling JSON file: %w", err)
	}
	return nil
}
