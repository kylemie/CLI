package main

import (
	"time"
)

type Bin struct {
	ID      string
	Private bool
	CreatAt time.Time
	Name    string
}

type BinList struct {
	Bins []Bin
}

func newBin(id string, private bool, creatat time.Time, name string) *Bin {
	b1 := Bin{
		ID:      id,
		Private: private,
		CreatAt: creatat,
		Name:    name,
	}
	return &b1
}

func newBinList(b *Bin) *BinList {
	bs := BinList{
		Bins: []Bin{*b},
	}
	return &bs
}

func main() {

}
