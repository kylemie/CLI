package file

import (
	"encoding/json"
	"os"
)

func ReadFile(name string) []byte {
	file, err := os.ReadFile(name)
	if err != nil {
		panic(err)
	}
	return file
}

func isJsonFile(file []byte) bool {
	var js json.RawMessage
	flag := json.Unmarshal(file, &js) == nil
	return flag
}

func WriteFile(file []byte, name string) {
	err := os.WriteFile(name, file, 0644)
	if err != nil {
		panic(err)
	}
}
