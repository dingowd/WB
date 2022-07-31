package model

type CacheOrder struct {
	TrackNumber       string   `json:"track_number" db:"track_number"`
	Entry             string   `json:"entry" db:"entry"`
	Delivery          Delivery `json:"delivery"`
	Payment           Payment  `json:"payment"`
	Items             []*Item  `json:"items"`
	Locale            string   `json:"locale" db:"locale"`
	InternalSignature string   `json:"internal_signature" db:"internal_signature"`
	CustomerId        string   `json:"customer_id" db:"customer_id"`
	DeliveryService   string   `json:"delivery_service" db:"delivery_service"`
	Shardkey          string   `json:"shardkey" db:"shardkey"`
	SmId              int      `json:"sm_id" db:"sm_id"`
	DateCreated       string   `json:"date_created" db:"date_created"`
	OofShard          string   `json:"oof_shard" db:"oof_shard"`
}

type CacheStruct map[string]CacheOrder
