package bins

import (
	"time"
)

type Bin struct {
	ID      string    `json:"id"`
	Private bool      `json:"private"`
	CreatAt time.Time `json:"creatat"`
	Name    string    `json:"name"`
}

type BinList struct {
	Bins []Bin `json:"bins"`
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
