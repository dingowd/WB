package postgres

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dingowd/WB/weather/service/internal/logger"
	"github.com/dingowd/WB/weather/service/models"
	"net/http"
	"sync"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"time"
)

type Storage struct {
	WG    sync.WaitGroup
	DB    *sqlx.DB
	Log   logger.Logger
	AppId string
}

func New(log logger.Logger, appid string) *Storage {
	return &Storage{Log: log, AppId: appid}
}

func (s *Storage) Connect(ctx context.Context, dsn string) error {
	var err error
	s.DB, err = sqlx.Open("pgx", dsn)
	if err == nil {
		s.Log.Info("База " + dsn + " подключена")
	} else {
		s.Log.Error("Ошибка соединения с базой. Проверьте параметры подключения")
	}
	return err
}

func (s *Storage) Close() error {
	s.Log.Info("Закрытие соединения с БД")
	return s.DB.Close()
}

func (s *Storage) GetCities() ([]models.City, error) {
	cities := make([]models.City, 0)
	query := `select city_id, name, lat, lon, country, state from cities order by name asc`
	rows, err := s.DB.QueryContext(context.Background(), query)
	if err != nil {
		return cities, err
	}
	defer rows.Close()
	for rows.Next() {
		var city models.City
		err := rows.Scan(&city.CityId, &city.Name, &city.Lat, &city.Lon, &city.Country, &city.State)
		if err != nil {
			continue
		}
		cities = append(cities, city)
	}
	return cities, nil

}

func (s *Storage) ShortWeather(city string) (models.ShortWeather, error) {
	var short models.ShortWeather
	cityToQuery := "'%" + city + "%'"
	query := `select detail, country from weather inner join cities on weather.city_id = cities.city_id where cities.name like ` + cityToQuery
	row := s.DB.QueryRow(query)
	var result, country string
	err := row.Scan(&result, &country)
	if err != nil {
		return short, err
	}
	var full models.Resp
	json.Unmarshal([]byte(result), &full)
	// вычисляем среднюю температуру
	var av, count float64
	for _, v := range full.List {
		av += (v.Main.TempMin + v.Main.TempMax) / 2
		count++
	}
	av = av / count
	// заполняем даты, на которые доступен прогноз
	date := ""
	for _, v := range full.List {
		t, _ := time.Parse("2006-01-02 15:04:05", v.DtTxt)
		ts := t.Format("02.01.2006")
		if ts != date {
			short.Dates = append(short.Dates, ts)
		}
		date = ts
	}
	short.Country = country
	short.City = city
	short.Date = time.Now().Format("02.01.2006")
	short.AvTemp = av
	return short, nil
}

func (s *Storage) DetWeather(city, date string) (models.Resp, error) {
	var resp, answer models.Resp
	_, err := time.Parse("02.01.2006", date)
	if err != nil {
		return answer, errors.New("Error. Please enter the date in dd.mm.yyyy format")
	}
	cityToQuery := "'%" + city + "%'"
	query := `select detail, country from weather inner join cities on weather.city_id = cities.city_id where cities.name like ` + cityToQuery
	row := s.DB.QueryRow(query)
	var result, country string
	err = row.Scan(&result, &country)
	if err != nil {
		return answer, err
	}
	json.Unmarshal([]byte(result), &resp)
	answer.Cod = resp.Cod
	answer.Message = resp.Message
	for _, v := range resp.List {
		t, _ := time.Parse("2006-01-02 15:04:05", v.DtTxt)
		if date == t.Format("02.01.2006") {
			answer.List = append(answer.List, v)
		}
	}
	answer.Cnt = len(answer.List)
	answer.City = resp.City
	return answer, nil
}

func (s *Storage) GetWeather() error {
	// получаем список городов из базы
	cities, err := s.GetCities()
	if err != nil {
		return err
	}
	// получаем погоду по API
	var mu sync.Mutex
	var wg sync.WaitGroup
	w := make([]models.Resp, 0)
	for _, v := range cities {
		wg.Add(1)
		go func(v models.City) {
			defer wg.Done()
			var elem models.Resp
			req := "http://api.openweathermap.org/data/2.5/forecast?lat=" +
				fmt.Sprint(v.Lat) + "&lon=" + fmt.Sprint(v.Lon) + "&lang=RU&units=metric&appid=" + s.AppId
			resp, err := http.Get(req)
			if err != nil {
				return
			}
			json.NewDecoder(resp.Body).Decode(&elem)
			defer resp.Body.Close()
			elem.CityId = v.CityId
			mu.Lock()
			w = append(w, elem)
			mu.Unlock()
		}(v)
	}
	wg.Wait()
	// чистим таблицу
	s.WG.Add(1)
	query := `truncate table weather`
	s.DB.Exec(query)
	// запись в бд
	t := time.Now().Format("02.01.2006")
	query = `insert into weather(city_id, date, temp, detail) values($1, $2, $3, $4)`
	for i := 0; i < len(cities); i++ {
		var count int
		var temp float64
		// вычисляем среднюю температуру на дату прогноза
		for _, v := range w[i].List {
			t2, _ := time.Parse("2006-01-02 15:04:05", v.DtTxt)
			if t == t2.Format("02.01.2006") {
				temp = (v.Main.TempMin + v.Main.TempMax) / 2
				count++
			}
		}
		temp = temp / float64(count)
		j, _ := json.Marshal(w[i])
		s.DB.ExecContext(context.Background(), query, w[i].CityId, t, temp, j)
	}
	s.WG.Done()
	return nil
}

func (s *Storage) Wait() {
	s.WG.Wait()
	s.Log.Info("GetWheather Done")
}

func (s *Storage) InsertUser(name string) error {
	query := `insert into favor(user_name) values($1)`
	_, err := s.DB.Exec(query, name)
	return err
}

func (s *Storage) InsertFav(name, city string) error {
	// Проверка существования пользователя в базе
	if !s.IsUserExist(name) {
		return errors.New("user is not exist")
	}
	// Проверка существования города в базе
	if !s.IsCityExist(city) {
		return errors.New("city is not exist")
	}
	query := `select favor from favor where user_name = $1`
	var cities []string
	if err := s.DB.QueryRow(query, name).Scan(pq.Array(&cities)); err != nil {
		return err
	}
	// Проверка города на входимость для исключения дупликации
	flag := false
	for _, v := range cities {
		if v == city {
			flag = true
			break
		}
	}
	if !flag {
		cities = append(cities, city)
	} else {
		return errors.New("city already in favorities")
	}
	// Обновляем избранные города
	query = `update favor set favor = $1 where user_name = $2`
	if _, err := s.DB.Exec(query, pq.Array(cities), name); err != nil {
		return err
	}
	return nil
}

func (s *Storage) GetFavor(name string) ([]string, error) {
	var res []string
	query := `select favor from favor where user_name = $1`
	err := s.DB.QueryRow(query, name).Scan(pq.Array(&res))
	return res, err
}

func (s *Storage) IsUserExist(name string) bool {
	query := `select user_name from favor where user_name = $1`
	var res string
	if err := s.DB.QueryRow(query, name).Scan(&res); err != nil {
		return false
	}
	return true
}

func (s *Storage) IsCityExist(city string) bool {
	query := `select name from cities where name = $1`
	var c string
	if err := s.DB.QueryRow(query, city).Scan(&c); err != nil {
		return false
	}
	return true
}
