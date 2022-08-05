package model

type CacheOrder struct {
	Order     Order
	TimeStamp int64
}

type CacheOrderList []CacheOrder
