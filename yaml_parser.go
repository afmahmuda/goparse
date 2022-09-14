package goparse

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

func parseYamlFile[T any](path string) (T, error) {
	path = filepath.Clean(path)

	var res T
	bs, err := os.ReadFile(path) // #nosec
	if err != nil {
		return res, err
	}

	if err := yaml.Unmarshal(bs, &res); err != nil {
		return res, err
	}

	return res, nil
}
