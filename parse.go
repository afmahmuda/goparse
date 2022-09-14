package goparse

import (
	"fmt"
	"path/filepath"
	"strings"
)

func ParseFile[T any](path string) (T, error) {
	path = filepath.Clean(path)

	var dafaultValue T

	ext := filepath.Ext(path)
	switch strings.ToLower(ext) {
	case ".yaml", ".yml":
		return parseYamlFile[T](path)
	case ".json":
		return parseJsonFile[T](path)
	}

	return dafaultValue, fmt.Errorf("extension unrecognizable `%s`", ext)
}
