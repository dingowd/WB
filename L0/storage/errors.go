package storage

import "errors"

var (
	ErrorOrderExist    error
	ErrorDeliveryExist error
)

func init() {
	ErrorOrderExist = errors.New("order already exist")
	ErrorDeliveryExist = errors.New("delivery already exist")
}
