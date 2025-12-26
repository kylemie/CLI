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
	err = json.Unmarshal(file, &name)
	if err != nil {
		panic(err)
	}
	return file
}

func WriteFile(file []byte, name string) {
	err := os.WriteFile(name, file, 0644)
	if err != nil {
		panic(err)
	}
}
