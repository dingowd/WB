package storage

import "errors"

var (
	ErrorDateBusy      error
	ErrorEventExist    error
	ErrorEventNotExist error
	DateFormatError    error
)

func init() {
	ErrorDateBusy = errors.New("event date busy")
	ErrorEventExist = errors.New("event already exist")
	ErrorEventNotExist = errors.New("event not exist")
	DateFormatError = errors.New("date format error")
}
