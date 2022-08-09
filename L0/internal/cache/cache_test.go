package cache

import (
	"github.com/dingowd/WB/L0/mocks"
	"github.com/dingowd/WB/L0/model"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"testing"
)

func TestCache(t *testing.T) {
	var a Cache
	l := mocks.NewLogger(t)
	stor := mocks.NewStorage(t)
	c := NewCache(l, stor, 25)
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

func TestInit(t *testing.T) {
	l := mocks.NewLogger(t)
	stor := mocks.NewStorage(t)
	c := NewCache(l, stor, 25)
	require.NotNil(t, c)
	stor.On("GetOrdersByLimit", mock.Anything).Once().Return(model.CacheOrderList{}, nil)
	c.Init()
	stor.AssertExpectations(t)
}

func TestReadFromCache(t *testing.T) {
	l := mocks.NewLogger(t)
	stor := mocks.NewStorage(t)
	c := NewCache(l, stor, 0)
	require.NotNil(t, c)
	stor.On("IsOrderExist", mock.Anything).Maybe().Return(true)
	stor.On("GetOrder", mock.Anything).Maybe().Return(model.Order{}, nil)
}
