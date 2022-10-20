package models

type City struct {
	CityId  int     `json:"-" db:"city_id"`
	Name    string  `json:"name" db:"name"`
	Lat     float64 `json:"-" db:"lat"`
	Lon     float64 `json:"-" db:"lon"`
	Country string  `json:"-" db:"country"`
	State   string  `json:"-" db:"state"`
}
