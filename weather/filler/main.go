package main

import (
	"encoding/json"
	"fmt"
	"github.com/BurntSushi/toml"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"net/http"
	"os"
)

type CityShort struct {
	Name    string `json:"name" db:"name"`
	Country string `json:"country" db:"country"`
	State   string `json:"state" db:"state"`
}
type CityShortList []CityShort

type City struct {
	Name    string  `json:"name" db:"name"`
	Lat     float64 `json:"lat" db:"lat"`
	Lon     float64 `json:"lon" db:"lon"`
	Country string  `json:"country" db:"country"`
	State   string  `json:"state" db:"state"`
}
type CityList []City

type Config struct {
	DSN   string
	File  string
	Appid string
}

func main() {
	conf := &Config{}
	cities := &CityShortList{}
	citiesToDB := make(CityList, 0)
	// Read config
	if _, err := toml.DecodeFile("config.toml", &conf); err != nil {
		fmt.Fprintln(os.Stdout, "ошибка чтения toml файла "+err.Error())
		return
	}
	// Read file of cities
	data, _ := os.ReadFile(conf.File)
	json.Unmarshal(data, cities)
	// Fill array to write to db
	for i := 0; i < len(*cities); i++ {
		url := "http://api.openweathermap.org/geo/1.0/direct?q=" + (*cities)[i].Name + "," +
			(*cities)[i].State + "," + (*cities)[i].Country + "&limit=5&appid=" + conf.Appid
		resp, err := http.Get(url)
		if err != nil {
			return
		}
		var c CityList
		json.NewDecoder(resp.Body).Decode(&c)
		citiesToDB = append(citiesToDB, c[0])
	}
	// Connect to db
	db, err := sqlx.Open("pgx", conf.DSN)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	// Write to db
	for _, v := range citiesToDB {
		query := `insert into cities(name, lat, lon, country, state) values($1, $2, $3, $4, $5)`
		_, err := db.Exec(query, v.Name, v.Lat, v.Lon, v.Country, v.State)
		if err != nil {
			fmt.Println(err)
		}
	}
}
