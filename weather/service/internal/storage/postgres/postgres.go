package postgres

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dingowd/WB/weather/service/internal/logger"
	"github.com/dingowd/WB/weather/service/models"
	"net/http"
	"sync"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"time"
)

type Storage struct {
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
	return short, nil

}

func (s *Storage) DetWeather(city string, t time.Time) (models.Resp, error) {
	var resp models.Resp
	return resp, nil

}

func (s *Storage) GetWeather() error {
	// чистим таблицу
	query := `truncate table weather`
	s.DB.Exec(query)
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
				fmt.Sprint(v.Lat) + "&lon=" + fmt.Sprint(v.Lon) + "&units=metric&appid=" + s.AppId
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
	// запись в бд
	t := time.Now().Format("02.01.2006")
	query = `insert into weather(city_id, date, temp, detail) values($1, $2, $3, $4)`
	for i := 0; i < len(cities); i++ {
		temp := w[i].List[1].Main.Temp
		j, _ := json.Marshal(w[i])
		s.DB.ExecContext(context.Background(), query, w[i].CityId, t, temp, j)
	}

	return nil
}
