package storage

import (
	"context"
	"github.com/dingowd/WB/weather/service/models"
)

type Storage interface {
	Connect(ctx context.Context, dsn string) error
	Close() error
	GetCities() ([]models.City, error)
	ShortWeather(city string) (models.ShortWeather, error)
	DetWeather(city, t string) (models.Resp, error)
	GetWeather() error
	Wait()
}
