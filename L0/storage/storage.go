package storage

import (
	"context"
	"github.com/dingowd/WB/L0/model"
)

type Storage interface {
	Connect(ctx context.Context, dsn string) error
	Close() error
	IsExist(id string) (bool, error)
	Create(d model.Order) error
	Get(id string) (model.Order, error)
}
