package api

import (
	"encoding/json"
)

type myBook struct {
	Title  string
	Author string
	ISBN   string
}

func (b myBook) ToJSON() []byte {
	ToJSON, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	return ToJSON
}
