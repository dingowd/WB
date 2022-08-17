package internalhttp

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func loggingMiddleware(f http.HandlerFunc, logg Logger) http.HandlerFunc {
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
