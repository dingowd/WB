package cache

import (
	"github.com/dingowd/WB/L0/model"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCache(t *testing.T) {
	var c, a Cache
	c.Amount = 25
	item := model.Item{
		ChrtId: 9934930,
	}
	items := make([]model.Item, 0)
	items = append(items, item)
	o := model.Order{
		OrderUid:    "b563feb7b2b84b6test",
		TrackNumber: "WBILMTESTTRACK",
		CustomerId:  "test",
		Items:       items,
	}
	c.Body = make(model.CacheOrderList, 0)
	c.WriteToCache(o)
	a.Body = c.Body
	require.Equal(t, o, a.Body[0].Order)
	var p model.CacheOrder
	b, _ := c.ReadFromCache("b563feb7b2b84b6test")
	p = *b
	require.Equal(t, o, p.Order)
}
