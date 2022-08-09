package utils

import (
	"encoding/json"
	"errors"
	"github.com/dingowd/WB/L0/model"
)

func Validate(data []byte) (*model.Order, error) {
	var o model.Order
	err := json.Unmarshal(data, &o)
	checkErr := errors.New("incorrect format")
	if err != nil {
		return nil, err
	}
	o1 := len(o.OrderUid) == 0
	o2 := len(o.TrackNumber) == 0
	o3 := len(o.Entry) == 0
	o4 := len(o.Locale) != 2
	o5 := len(o.CustomerId) == 0
	o6 := len(o.DeliveryService) == 0
	o7 := len(o.Shardkey) == 0
	o8 := o.SmId == 0
	o9 := len(o.DateCreated) == 0
	o10 := len(o.OofShard) == 0
	order := o1 || o2 || o3 || o4 || o5 || o6 || o7 || o8 || o9 || o10

	d1 := len(o.Delivery.Name) == 0
	d2 := len(o.Delivery.Phone) == 0
	d3 := len(o.Delivery.Zip) == 0
	d4 := len(o.Delivery.City) == 0
	d5 := len(o.Delivery.Address) == 0
	d6 := len(o.Delivery.Region) == 0
	d7 := len(o.Delivery.Email) == 0
	del := d1 || d2 || d3 || d4 || d5 || d6 || d7

	p1 := len(o.Payment.Transaction) == 0 || o.Payment.Transaction != o.OrderUid
	p2 := len(o.Payment.Currency) != 3
	p3 := len(o.Payment.Provider) == 0
	p4 := o.Payment.Amount == 0
	p5 := o.Payment.PaymentDt == 0
	p6 := len(o.Payment.Bank) == 0
	p7 := o.Payment.DeliveryCost == 0
	p8 := o.Payment.GoodsTotal == 0
	pay := p1 || p2 || p3 || p4 || p5 || p6 || p7 || p8

	itemsL := len(o.Items) == 0
	i := false
	for _, v := range o.Items {
		if v.ChrtId == 0 || len(v.TrackNumber) == 0 || v.Price == 0 || len(v.Rid) == 0 || len(v.Name) == 0 || len(v.Size) == 0 || v.TotalPrice == 0 || v.NmId == 0 || v.Status == 0 {
			i = true
			break
		}
	}

	if order || pay || del || itemsL || i {
		return nil, checkErr
	}
	return &o, nil
}
