package utils

import (
	"encoding/json"
	"github.com/dingowd/WB/L2/develop/dev11/models"
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

type ResArr struct {
	Res []models.DBEvent `json:"result"`
}

func ReturnResultArr(a []models.DBEvent) []byte {
	r := ResArr{
		Res: a,
	}
	b, _ := json.Marshal(r)
	return b
}
