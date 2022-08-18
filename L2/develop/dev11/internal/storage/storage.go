package storage

import (
	"context"
	"github.com/dingowd/WB/L2/develop/dev11/models"
)

type Storage interface {
	Connect(ctx context.Context, dsn string) error
	Close() error
	Create(e models.Event) error
	Update(id int, e models.Event) error
	Delete(id int) error
	GetDayEvent(day string) ([]models.Event, error)
	GetWeekEvent(day string) ([]models.Event, error)
	GetMonthEvent(day string) ([]models.Event, error)
}
