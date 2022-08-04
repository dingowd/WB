package postgres

import (
	"context"
	"database/sql"
	"github.com/dingowd/WB/L0/logger"
	"github.com/dingowd/WB/L0/model"
	"github.com/dingowd/WB/L0/storage"
	"github.com/jmoiron/sqlx"

	_ "github.com/jackc/pgx/stdlib"
)

type Storage struct {
	DB  *sqlx.DB
	Log logger.Logger
}

func New(log logger.Logger) *Storage {
	return &Storage{Log: log}
}

func (s *Storage) Connect(ctx context.Context, dsn string) error {
	var err error
	s.DB, err = sqlx.Open("pgx", dsn)
	if err == nil {
		s.Log.Info("Connected to " + dsn)
	} else {
		s.Log.Error("Failed to connect database")
	}
	return err
}

func (s *Storage) Close() error {
	return s.DB.Close()
}

func (s *Storage) CreateOrder(d model.Order) error {
	msg := "Creating order with ID " + d.OrderUid
	s.Log.Info(msg)
	exist, _ := s.IsOrderExist(d.OrderUid)
	if exist {
		msg = "Order with ID " + d.OrderUid + " already exist. Order creation error."
		s.Log.Info(msg)
		return storage.ErrorOrderExist
	}
	var deliveryID int
	if exist, deliveryID, _ = s.IsDeliveryExist(d.Delivery); !exist {
		deliveryID, _ = s.CreateDelivery(d.Delivery)
	}
	err := s.CreatePayment(d.Payment)
	query := "insert into orders (order_uid, track_number, entry, locale, internal_signature, customer_id, delivery_service, " +
		"shardkey, sm_id, date_created, oof_shard, delivery_id, transaction) " +
		"values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)"
	_, err = s.DB.Exec(query, d.OrderUid, d.TrackNumber, d.Entry, d.Locale, d.InternalSignature, d.CustomerId, d.DeliveryService,
		d.Shardkey, d.SmId, d.DateCreated, d.OofShard, deliveryID, d.Payment.Transaction)
	err = s.CreateItems(d.OrderUid, d.Items)
	if err != nil {
		return err
	}
	msg = "Successful creation of the order with ID " + d.OrderUid
	s.Log.Info(msg)
	return nil
}

func (s *Storage) GetOrder(id string) (*model.Order, error) {
	query := "select order_uid, track_number, entry, " +
		"name, phone, zip, city, address, region, email, " +
		"request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee, " +
		"locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, " +
		"oof_shard, delivery_id, orders.transaction as transaction " +
		"from orders " +
		"inner join delivery on orders.delivery_id = delivery.id " +
		"inner join payment on orders.transaction = payment.transaction " +
		"where order_uid = $1"
	var order model.Order
	rows, err := s.DB.NamedQuery(query, id)
	if err != nil {
		msg := "Error to get order with ID " + err.Error()
		s.Log.Error(msg)
		return nil, err
	}
	fromDB := make([]model.DbOrderNoItems, 0)
	for rows.Next() {
		var elem model.DbOrderNoItems
		err := rows.StructScan(&elem)
		if err != nil {
			msg := "Error to get order with ID " + err.Error()
			s.Log.Error(msg)
			return nil, err
		}
		fromDB = append(fromDB, elem)
	}
	e := fromDB[0]
	order.OrderUid = e.OrderUid
	order.TrackNumber = e.TrackNumber
	order.Entry = e.Entry
	order.Delivery.Name = e.Name
	order.Delivery.Phone = e.Phone
	order.Delivery.Zip = e.Zip
	order.Delivery.City = e.City
	order.Delivery.Address = e.Address
	order.Delivery.Region = e.Region
	order.Delivery.Email = e.Email
	order.Payment.Transaction = e.Transaction
	order.Payment.RequestId = e.RequestId
	order.Payment.Currency = e.Currency
	order.Payment.Amount = e.Amount
	order.Payment.PaymentDt = e.PaymentDt
	order.Payment.Bank = e.Bank
	order.Payment.DeliveryCost = e.DeliveryCost
	order.Payment.GoodsTotal = e.GoodsTotal
	order.Payment.CustomFee = e.CustomFee
	order.Locale = e.Locale
	order.InternalSignature = e.InternalSignature
	order.CustomerId = e.CustomerId
	order.DeliveryService = e.DeliveryService
	order.Shardkey = e.Shardkey
	order.SmId = e.SmId
	order.DateCreated = e.DateCreated
	order.OofShard = e.OofShard
	order.Items, err = s.GetItems(order.OrderUid)
	return &order, nil
}

func (s *Storage) IsOrderExist(id string) (bool, error) {
	query := "select order_uid from orders where order_uid = $1"
	row := s.DB.QueryRow(query, id)
	var orderId string
	err := row.Scan(&orderId)
	if err == sql.ErrNoRows {
		return false, nil
	} else if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

func (s *Storage) IsDeliveryExist(d model.Delivery) (bool, int, error) {
	query := "select id from delivery where name = $1 and phone = $2 and zip = $3 and city = $4 and address = $5 and region = $6 and email = $7"
	row := s.DB.QueryRow(query, d.Name, d.Phone, d.Zip, d.City, d.Address, d.Region, d.Email)
	var orderId int
	err := row.Scan(&orderId)
	if err == sql.ErrNoRows {
		return false, 0, nil
	} else if err != nil {
		return false, 0, err
	} else {
		return true, orderId, nil
	}
}

func (s *Storage) CreateDelivery(d model.Delivery) (int, error) {
	var err error
	var id int
	query := "insert into delivery (name, phone, zip, city, address, region, email) values ($1, $2, $3, $4, $5, $6, $7)"
	_, err = s.DB.Exec(query, d.Name, d.Phone, d.Zip, d.City, d.Address, d.Region, d.Email)
	if err != nil {
		return 0, err
	}
	_, id, err = s.IsDeliveryExist(d)
	return id, err
}

func (s *Storage) CreatePayment(p model.Payment) error {
	query := "insert into payment (transaction, request_id, currency, provider, amount, payment_dt, bank, " +
		"delivery_cost, goods_total, custom_fee) " +
		"values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)"
	_, err := s.DB.Exec(query, p.Transaction, p.RequestId, p.Currency, p.Provider, p.Amount, p.PaymentDt, p.Bank,
		p.DeliveryCost, p.GoodsTotal, p.CustomFee)
	return err
}

func (s *Storage) CreateItems(id string, i []model.Item) error {
	var err error
	query := "insert into items (chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status, order_uid) " +
		"values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)"
	for _, val := range i {
		if !s.IsItemExist(id, val) {
			_, err = s.DB.Exec(query, val.ChrtId, val.TrackNumber, val.Price, val.Rid, val.Name, val.Sale, val.Size,
				val.TotalPrice, val.NmId, val.Brand, val.Status, id)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *Storage) IsItemExist(id string, i model.Item) bool {
	query := "select chrt_id from items where chrt_id = $1 and track_number = $2 and price = $3 and rid = $4 and name = $5 " +
		"and sale = $6 and size = $7 and total_price = $8 and nm_id = $9 and brand = $10 and status = $11 and order_uid = $12"
	row := s.DB.QueryRow(query, i.ChrtId, i.TrackNumber, i.Price, i.Rid, i.Name, i.Sale, i.Size, i.TotalPrice, i.NmId, i.Brand,
		i.Status, id)
	var orderId int
	err := row.Scan(&orderId)
	if err == sql.ErrNoRows {
		return false
	} else if err != nil {
		return false
	} else {
		return true
	}
}

func (s *Storage) GetItems(id string) ([]model.Item, error) {
	query := "select chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status " +
		"from items where id = $1"
	items := make([]model.Item, 0)
	rows, err := s.DB.NamedQuery(query, id)
	if err != nil {
		return items, err
	}
	for rows.Next() {
		var item model.Item
		err := rows.StructScan(&item)
		if err != nil {
			return items, err
		}
		items = append(items, item)
	}
	return items, nil
}
