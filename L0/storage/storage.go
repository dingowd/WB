package storage

import (
	"context"
	"github.com/dingowd/WB/L0/model"
)

type Storage interface {
	Connect(ctx context.Context, dsn string) error
	Close() error
	CreateOrder(d model.Order) error
	GetOrder(id string) (model.Order, error)
	GetOrdersByLimit(a int) (model.CacheOrderList, error)
}
