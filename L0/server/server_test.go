package server

import (
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServerStatus(t *testing.T) {
	var s Server
	r := httptest.NewRequest("GET", "http://127.0.0.1:3541/get", nil)
	w := httptest.NewRecorder()
	s.GetOrder(w, r)
	resp := w.Result()
	require.Equal(t, http.StatusBadRequest, resp.StatusCode)
}
