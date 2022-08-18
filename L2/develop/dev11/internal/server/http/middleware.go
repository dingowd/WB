package internalhttp

import (
	"encoding/json"
	"github.com/dingowd/WB/L2/develop/dev11/internal/logger"
	"io/ioutil"
	"net/http"
)

func loggingMiddleware(f http.HandlerFunc, logg logger.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			logg.Error(err.Error())
			return
		}
		var s string
		if err := json.Unmarshal(content, &s); err != nil {
			logg.Error(err.Error())
			return
		}
		logg.Info(s)
		f(w, r)
	}
}
