package storage

import (
	"context"
	"github.com/dingowd/WB/weather/service/models"
	"time"
)

type Storage interface {
	Connect(ctx context.Context, dsn string) error
	Close() error
	GetCities() ([]models.City, error)
	ShortWeather(city string) (models.ShortWeather, error)
	DetWeather(city string, t time.Time) (models.Resp, error)
	GetWeather() error
}
