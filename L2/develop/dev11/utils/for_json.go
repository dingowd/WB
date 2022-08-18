package utils

import (
	"encoding/json"
	"github.com/dingowd/WB/L2/develop/dev11/models"
	"net/http"
)

func ToStruct(r *http.Request, e *models.Event) error {
	return json.NewDecoder(r.Body).Decode(e)
}

func ToDBStruct(r *http.Request, e *models.DBEvent) error {
	return json.NewDecoder(r.Body).Decode(e)
}
