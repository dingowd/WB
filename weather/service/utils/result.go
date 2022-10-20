package utils

import (
	"encoding/json"
)

type Err struct {
	Err string `json:"error"`
}

func ReturnError(s string) []byte {
	e := Err{
		Err: s,
	}
	b, _ := json.Marshal(e)
	return b
}

type Res struct {
	Res string `json:"result"`
}

func ReturnResult(s string) []byte {
	r := Res{
		Res: s,
	}
	b, _ := json.Marshal(r)
	return b
}
