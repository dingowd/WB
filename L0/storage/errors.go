package storage

import "errors"

var (
	ErrorExist error
)

func init() {
	ErrorExist = errors.New("order already exist")
}
