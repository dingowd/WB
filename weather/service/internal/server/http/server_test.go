package internalhttp

import (
	"context"
	"github.com/dingowd/WB/weather/service/internal/app"
	"github.com/dingowd/WB/weather/service/internal/logger"
	"github.com/dingowd/WB/weather/service/internal/storage"
	"github.com/dingowd/WB/weather/service/models"
	"github.com/go-chi/chi"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockStorage struct {
}

func (m *MockStorage) Connect(ctx context.Context, dsn string) error {
	return nil
}

func (m *MockStorage) Close() error {
	return nil
}

func (m *MockStorage) GetCities() ([]models.City, error) {
	e := make([]models.City, 0)
	return e, nil
}

func (m *MockStorage) ShortWeather(city string) (models.ShortWeather, error) {
	var e models.ShortWeather
	return e, nil
}

func (m *MockStorage) DetWeather(city, t string) (models.Resp, error) {
	var e models.Resp
	return e, nil
}

func (m *MockStorage) GetWeather() error {
	return nil
}

func (m *MockStorage) Wait() {}

func (m *MockStorage) InsertUser(name string) error {
	return nil
}

func (m *MockStorage) InsertFav(name, city string) error {
	return nil
}

func (m *MockStorage) GetFavor(name string) ([]string, error) {
	e := make([]string, 0)
	return e, nil
}

func TestGetCities(t *testing.T) {
	app := app.New(*new(logger.Logger), *new(storage.Storage))
	var m MockStorage
	app.Storage = &m
	s := NewServer(app, "127.0.0.1:3541")
	router := chi.NewRouter()
	s.Srv = &http.Server{Addr: s.Addr, Handler: router}
	r := httptest.NewRequest("GET", "http://127.0.0.1:3541/cities", nil)
	w := httptest.NewRecorder()
	s.GetCities(w, r)
	resp := w.Result()
	require.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestGetShort(t *testing.T) {
	app := app.New(*new(logger.Logger), *new(storage.Storage))
	var m MockStorage
	app.Storage = &m
	s := NewServer(app, "127.0.0.1:3541")
	router := chi.NewRouter()
	s.Srv = &http.Server{Addr: s.Addr, Handler: router}
	r1 := httptest.NewRequest("GET", "http://127.0.0.1:3541/short", nil)
	w1 := httptest.NewRecorder()
	s.GetShort(w1, r1)
	resp := w1.Result()
	require.Equal(t, http.StatusBadRequest, resp.StatusCode)

	r2 := httptest.NewRequest("GET", "http://127.0.0.1:3541/short?city=Moscow", nil)
	w2 := httptest.NewRecorder()
	s.GetShort(w2, r2)
	resp = w2.Result()
	require.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestGetDetail(t *testing.T) {
	app := app.New(*new(logger.Logger), *new(storage.Storage))
	var m MockStorage
	app.Storage = &m
	s := NewServer(app, "127.0.0.1:3541")
	router := chi.NewRouter()
	s.Srv = &http.Server{Addr: s.Addr, Handler: router}
	r := httptest.NewRequest("GET", "http://127.0.0.1:3541/detail?city=Moscow&date=26.10.2022", nil)
	w := httptest.NewRecorder()
	s.GetDetail(w, r)
	resp := w.Result()
	require.Equal(t, http.StatusOK, resp.StatusCode)
}
