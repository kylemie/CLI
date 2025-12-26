package storage

import (
	"cli/pet-project/bins"
	"cli/pet-project/file"
	"encoding/json"
)

func ReadBinJSON(name string) *bins.BinList {
	d := file.ReadFile(name)
	var data bins.BinList
	_ = json.Unmarshal(d, &data)
	return &data
}

func SaveBin(bin *bins.BinList, name string) {
	b, err := json.Marshal(bin)
	if err != nil {
		panic(err)
	}
	file.WriteFile(b, name)
}
