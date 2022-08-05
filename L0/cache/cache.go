package cache

import (
	"github.com/dingowd/WB/L0/logger"
	"github.com/dingowd/WB/L0/model"
	"github.com/dingowd/WB/L0/storage"
	"time"
)

type CacheInterface interface {
	Init()
	ReadFromCache(id string) model.CacheOrder
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
	var err error
	//c.Body = make(model.CacheOrderList, 0)
	c.Body, err = c.Stor.GetOrdersByLimit(c.Amount)
	if err != nil {
		c.Log.Error("Unable to fill cache " + err.Error())
	}
}

func (c *Cache) ReadFromCache(id string) model.CacheOrder {
	for i := 0; i < len(c.Body); i++ {
		if id == c.Body[i].Order.OrderUid {
			c.Body[i].TimeStamp = time.Now().UnixNano()
			return c.Body[i]
		}
	}
	min := c.Body[0].TimeStamp
	key := 0
	for i := 1; i < len(c.Body); i++ {
		if c.Body[i].TimeStamp < min {
			min = c.Body[i].TimeStamp
			key = i
		}
	}
	c.Body[key].Order, _ = c.Stor.GetOrder(id)
	c.Body[key].TimeStamp = time.Now().UnixNano()
	return c.Body[key]
}
