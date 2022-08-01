package postgres

import (
	"context"
	"fmt"
	"github.com/dingowd/WB/L0/logger"
	"github.com/dingowd/WB/L0/model"
	"github.com/dingowd/WB/L0/storage"
	"github.com/jmoiron/sqlx"
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
	s.DB, err = sqlx.Open("postgres", dsn)
	return err
}

func (s *Storage) Close() error {
	return s.DB.Close()
}

func (s *Storage) Create(d model.Order) error {
	exist, _ := s.IsExist(d.OrderUid)
	if exist {
		return storage.ErrorExist
	}
	// TODO

	return nil
}

func (s *Storage) Get(id string) (model.Order, error) {
	order := &model.Order{}

}

func (s *Storage) IsExist(id string) (bool, error) {
	rows, err := s.DB.Query("select order_uid from orders where order_uid = $1", id)
	if err != nil {
		return false, err
	}
	if rows.Err() != nil {
		return false, rows.Err()
	}
	defer rows.Close()
	ids := make([]string, 0)
	for rows.Next() {
		var id string
		err := rows.Scan(&id)
		if err != nil {
			s.Log.Error(err.Error())
			continue
		}
		ids = append(ids, id)
	}
	return len(ids) != 0, nil
}
