package models

type City struct {
	CityId  int     `json:"city_id" db:"city_id"`
	Name    string  `json:"name" db:"name"`
	Lat     float64 `json:"lat" db:"lat"`
	Lon     float64 `json:"lon" db:"lon"`
	Country string  `json:"country" db:"country"`
	State   string  `json:"state" db:"state"`
}
