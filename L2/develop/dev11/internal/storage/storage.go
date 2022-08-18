package storage

import (
	"context"
	"github.com/dingowd/WB/L2/develop/dev11/models"
)

type Storage interface {
	Connect(ctx context.Context, dsn string) error
	Close() error
	Create(e models.Event) error
	Update(e models.DBEvent) error
	Delete(id int) error
	GetDayEvent(id int, day string) ([]models.DBEvent, error)
	GetWeekEvent(id int, day string) ([]models.DBEvent, error)
	GetMonthEvent(id int, day string) ([]models.DBEvent, error)
}
