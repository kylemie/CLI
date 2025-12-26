package file

import (
	"os"
	"path/filepath"
	"strings"
)

func ReadFile(name string) []byte {
	file, err := os.ReadFile(name)
	if err != nil {
		panic(err)
	}
	return file
}

func IsJsonFile(name string) bool {
	ext := filepath.Ext(name)
	return strings.ToLower(ext) == ".json"
}

func WriteFile(file []byte, name string) {
	err := os.WriteFile(name, file, 0644)
	if err != nil {
		panic(err)
	}
}
