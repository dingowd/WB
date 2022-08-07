package cache

import (
	"github.com/dingowd/WB/L0/logger"
	"github.com/dingowd/WB/L0/model"
	"github.com/dingowd/WB/L0/storage"
	"time"
)

type CacheInterface interface {
	Init()
	ReadFromCache(id string) (*model.CacheOrder, error)
}

type Cache struct {
	Log    logger.Logger
	Stor   storage.Storage
	Amount int
	Body   model.CacheOrderList
}

func NewCache(log logger.Logger, stor storage.Storage, a int) *Cache {
	return &Cache{
		Log:    log,
		Stor:   stor,
		Amount: a,
		Body:   make(model.CacheOrderList, 0),
	}
}

func (c *Cache) Init() {
	c.Body, _ = c.Stor.GetOrdersByLimit(c.Amount)
}

func (c *Cache) ReadFromCache(id string) (*model.CacheOrder, error) {
	for i := 0; i < len(c.Body); i++ {
		if id == c.Body[i].Order.OrderUid {
			c.Body[i].TimeStamp = time.Now().UnixNano()
			return &c.Body[i], nil
		}
	}
	var err error
	key := 0
	if len(c.Body) == c.Amount {
		min := c.Body[0].TimeStamp
		for i := 1; i < len(c.Body); i++ {
			if c.Body[i].TimeStamp < min {
				min = c.Body[i].TimeStamp
				key = i
			}
		}
		c.Body[key].Order, err = c.Stor.GetOrder(id)
		if err != nil {
			return nil, err
		}
		c.Body[key].TimeStamp = time.Now().UnixNano()
		return &c.Body[key], nil
	} else {
		var o model.Order
		var b model.CacheOrder
		if !c.Stor.IsOrderExist(id) {
			return nil, storage.ErrorOrderNotExist
		}
		o, err = c.Stor.GetOrder(id)
		if err != nil {
			return nil, err
		}
		b.Order = o
		b.TimeStamp = time.Now().UnixNano()
		c.Body = append(c.Body, b)
	}
	return &c.Body[len(c.Body)-1], nil
}
