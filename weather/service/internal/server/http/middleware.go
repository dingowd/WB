package internalhttp

import (
	"github.com/dingowd/WB/weather/service/internal/logger"
	"net/http"
)

func loggingMiddleware(f http.HandlerFunc, logg logger.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s := r.Method + " " + r.RequestURI
		logg.Info(s)
		f(w, r)
	}
}
