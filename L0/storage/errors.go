package storage

import "errors"

var (
	ErrorOrderCreate    error
	ErrorOrderExist     error
	ErrorPaymentExist   error
	ErrorDeliveryCreate error
	ErrorPaymentCreate  error
	ErrorItemIDExist    error
)

func init() {
	ErrorOrderCreate = errors.New("ошибка создания заказа")
	ErrorOrderExist = errors.New("заказ уже существует")
	ErrorPaymentExist = errors.New("оплата уже существует")
	ErrorDeliveryCreate = errors.New("ошибка создания доставки")
	ErrorPaymentCreate = errors.New("ошибка создания оплаты")
	ErrorItemIDExist = errors.New("ошибка. товар с таким ID уже существует")
}
