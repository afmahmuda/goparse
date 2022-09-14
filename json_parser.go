package goparse

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func parseJsonFile[T any](path string) (T, error) {
	path = filepath.Clean(path)

	var res T
	bs, err := os.ReadFile(path) // #nosec
	if err != nil {
		return res, err
	}

	if err := json.Unmarshal(bs, &res); err != nil {
		return res, err
	}

	return res, nil
}
