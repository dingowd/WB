package models

type Favor struct {
	UserName string   `json:"user_name" db:"user_name"`
	Favor    []string `json:"favor" db:"favor"`
}
