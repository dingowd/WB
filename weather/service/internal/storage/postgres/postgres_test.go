package postgres

import (
	"github.com/dingowd/WB/L0/internal/logger/lrus"
	"github.com/dingowd/WB/L0/model"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetOrder(t *testing.T) {
	var s Storage
	s.DB, _ = sqlx.Open("pgx", "user=postgres dbname=WB sslmode=disable password=masterkey")
	defer s.DB.Close()
	s.Log = lrus.New()
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
	err := s.CreateOrder(o)
	require.Error(t, err)
}

func TestIsOrderExist(t *testing.T) {
	var s Storage
	s.DB, _ = sqlx.Open("pgx", "user=postgres dbname=WB sslmode=disable password=masterkey")
	defer s.DB.Close()
	s.Log = lrus.New()
	ok := s.IsOrderExist("b563feb7b2b84b6test")
	require.Equal(t, true, ok)
}
