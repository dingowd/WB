package internalhttp

import (
	"encoding/json"
	"errors"
	_ "github.com/dingowd/WB/weather/service/docs"
	"github.com/dingowd/WB/weather/service/internal/app"
	"github.com/dingowd/WB/weather/service/models"
	"github.com/dingowd/WB/weather/service/utils"
	"github.com/go-chi/chi"
	httpSwagger "github.com/swaggo/http-swagger"
	"html/template"

	"net/http"
)

type Server struct {
	App  *app.App
	Addr string
	Srv  *http.Server
}

func NewServer(app *app.App, addr string) *Server {
	return &Server{App: app, Addr: addr}
}

var (
	ErrorStopServer  = errors.New("timeout to stop server")
	ErrorStartServer = errors.New("timeout to start server")
)

type Response struct {
	Msg string `json:"msg"`
}

type Request struct {
	Msg string `json:"msg"`
}

// GetCities godoc
// @Summary Список городов
// @Description Получить список городов
// @Produce json
// @Success 200 {object} models.City
// @Failure 400 {object} utils.Err
// @Failure 500 {object} utils.Err
// @Router /cities [get]
func (s *Server) GetCities(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(utils.ReturnError("Method isn`t GET"))
		return
	}
	// получаем список городов из базы
	cities, err := s.App.Storage.GetCities()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(utils.ReturnError("Error getting list of cities"))
		return
	}
	tmpl, errT := template.ParseFiles("./templates/cities.html")
	if errT != nil {
		b, err := json.Marshal(cities)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(utils.ReturnError("Error getting list of cities"))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	} else {
		tmpl.Execute(w, cities)
	}
}

// GetShort godoc
// @Summary Краткий прогноз
// @Description Получить краткий прогноз
// @Produce json
// @Param city query string true "Название города"
// @Success 200 {object} models.ShortWeather
// @Failure 400 {object} utils.Err
// @Failure 500 {object} utils.Err
// @Router /short [get]
func (s *Server) GetShort(w http.ResponseWriter, r *http.Request) {
	s.App.Storage.Wait()
	if r.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(utils.ReturnError("Method isn`t GET"))
		return
	}
	name := r.URL.Query().Get("city")
	if len(name) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(utils.ReturnError("Please enter the city name"))
		return
	}
	short, err := s.App.Storage.ShortWeather(name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(utils.ReturnError("Error getting weather in " + name + ". " + err.Error()))
		return
	}
	tmpl, errT := template.ParseFiles("./templates/short.html")
	if errT != nil {
		b, err := json.Marshal(short)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(utils.ReturnError("Error getting weather in " + name + ". " + err.Error()))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	} else {
		tmpl.Execute(w, short)
	}
}

// GetDetail godoc
// @Summary Детальный прогноз
// @Description Получить детальный прогноз
// @Produce json
// @Param city query string true "Название города"
// @Param date query string true "Дата прогноза"
// @Success 200 {object} models.Resp
// @Failure 400 {object} utils.Err
// @Failure 500 {object} utils.Err
// @Router /detail [get]
func (s *Server) GetDetail(w http.ResponseWriter, r *http.Request) {
	s.App.Storage.Wait()
	if r.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(utils.ReturnError("Method isn`t GET"))
		return
	}
	name := r.URL.Query().Get("city")
	date := r.URL.Query().Get("date")
	if len(name) == 0 || len(date) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(utils.ReturnError("Please enter the city name and date"))
		return
	}
	result, err := s.App.Storage.DetWeather(name, date)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(utils.ReturnError("Error getting weather in " + name + ". " + err.Error()))
		return
	}
	tmpl, errT := template.ParseFiles("./templates/detail.html")
	if errT != nil {
		b, err := json.Marshal(result)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(utils.ReturnError("Error getting weather in " + name + ". " + err.Error()))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	} else {
		var resWithDate models.RespWithDate
		resWithDate.Date = date
		resWithDate.City = result.City
		resWithDate.List = append(resWithDate.List, result.List...)
		tmpl.Execute(w, resWithDate)
	}
}

// InsertUser godoc
// @Summary Добавить пользователя
// @Description Добавить нового пользователя
// @Produce json
// @Param name query string true "Имя нового пользователя"
// @Success 200 {object} utils.Res
// @Failure 400 {object} utils.Err
// @Failure 500 {object} utils.Err
// @Router /insert_user [post]
func (s *Server) InsertUser(w http.ResponseWriter, r *http.Request) {
	s.App.Storage.Wait()
	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(utils.ReturnError("Method isn`t POST"))
		return
	}
	name := r.URL.Query().Get("name")
	if len(name) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(utils.ReturnError("Missing username"))
		return
	}
	err := s.App.Storage.InsertUser(name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(utils.ReturnError(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(utils.ReturnResult("User " + name + " added"))
}

// InsertFav godoc
// @Summary Добавить избранное
// @Description Добавить город в избранное
// @Produce json
// @Param name query string true "Имя пользователя"
// @Param city query string true "Название нового города"
// @Success 200 {object} utils.Res
// @Failure 400 {object} utils.Err
// @Failure 500 {object} utils.Err
// @Router /insert_fav [post]
func (s *Server) InsertFav(w http.ResponseWriter, r *http.Request) {
	s.App.Storage.Wait()
	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(utils.ReturnError("Method isn`t POST"))
		return
	}
	name := r.URL.Query().Get("name")
	city := r.URL.Query().Get("city")
	if len(name) == 0 || len(city) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(utils.ReturnError("Missing username or city"))
		return
	}
	err := s.App.Storage.InsertFav(name, city)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(utils.ReturnError(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(utils.ReturnResult("City " + city + " added"))
}

// GetShortFavor godoc
// @Summary Краткий прогноз
// @Description Получить краткий прогноз избранных городов
// @Produce json
// @Param name query string true "Имя пользователя"
// @Param city query string true "Название города"
// @Success 200 {array} models.ShortWeather
// @Failure 400 {object} utils.Err
// @Failure 500 {object} utils.Err
// @Router /short_favor [get]
func (s *Server) GetShortFavor(w http.ResponseWriter, r *http.Request) {
	s.App.Storage.Wait()
	if r.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(utils.ReturnError("Method isn`t GET"))
		return
	}
	name := r.URL.Query().Get("name")
	if len(name) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(utils.ReturnError("Please enter the user name"))
		return
	}
	favors, err := s.App.Storage.GetFavor(name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(utils.ReturnError("Error getting favorites of user " + name + ". " + err.Error()))
		return
	}
	shorts := make([]models.ShortWeather, 0)
	for _, v := range favors {
		short, err := s.App.Storage.ShortWeather(v)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(utils.ReturnError("Error getting weather in " + name + ". " + err.Error()))
			return
		}
		shorts = append(shorts, short)
	}
	tmpl, errT := template.ParseFiles("./templates/shorts.html")
	if errT != nil {
		b, err := json.Marshal(shorts)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(utils.ReturnError("Error getting favorities of user " + name + ". " + err.Error()))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	} else {
		tmpl.Execute(w, shorts)
	}
}

// GetDetailFavor godoc
// @Summary Детальный прогноз
// @Description Получить детальный прогноз избранных городов
// @Produce json
// @Param name query string true "Имя пользователя"
// @Param date query string true "Дата прогноза"
// @Success 200 {array} models.Resp
// @Failure 400 {object} utils.Err
// @Failure 500 {object} utils.Err
// @Router /detail_favor [get]
func (s *Server) GetDetailFavor(w http.ResponseWriter, r *http.Request) {
	s.App.Storage.Wait()
	if r.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(utils.ReturnError("Method isn`t GET"))
		return
	}
	name := r.URL.Query().Get("name")
	date := r.URL.Query().Get("date")
	if len(name) == 0 || len(date) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(utils.ReturnError("Please enter the city name and date"))
		return
	}
	favors, err := s.App.Storage.GetFavor(name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(utils.ReturnError("Error getting favorities of user " + name + ". " + err.Error()))
		return
	}
	details := make([]models.Resp, 0)
	for _, v := range favors {
		result, err := s.App.Storage.DetWeather(v, date)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(utils.ReturnError("Error getting weather in " + name + ". " + err.Error()))
			return
		}
		details = append(details, result)
	}
	tmpl, errT := template.ParseFiles("./templates/details.html")
	if errT != nil {
		b, err := json.Marshal(details)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(utils.ReturnError("Error getting weather in " + name + ". " + err.Error()))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	} else {
		results := make([]models.RespWithDate, 0)
		for _, v := range details {
			var resWithDate models.RespWithDate
			resWithDate.Date = date
			resWithDate.City = v.City
			resWithDate.List = append(resWithDate.List, v.List...)
			results = append(results, resWithDate)
		}
		tmpl.Execute(w, results)
	}
}

func (s *Server) Start() error {
	r := chi.NewRouter()

	r.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:1323/swagger/doc.json")))

	go http.ListenAndServe(":1323", r)

	s.App.Logg.Info("http server starting")
	mux := http.NewServeMux()
	s.Srv = &http.Server{Addr: s.Addr, Handler: mux}
	mux.HandleFunc("/cities", loggingMiddleware(s.GetCities, s.App.Logg))
	mux.HandleFunc("/short", loggingMiddleware(s.GetShort, s.App.Logg))
	mux.HandleFunc("/detail", loggingMiddleware(s.GetDetail, s.App.Logg))
	mux.HandleFunc("/insert_user", loggingMiddleware(s.InsertUser, s.App.Logg))
	mux.HandleFunc("/insert_fav", loggingMiddleware(s.InsertFav, s.App.Logg))
	mux.HandleFunc("/short_favor", loggingMiddleware(s.GetShortFavor, s.App.Logg))
	mux.HandleFunc("/detail_favor", loggingMiddleware(s.GetDetailFavor, s.App.Logg))
	//mux.HandleFunc("/swagger/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:3541/swagger/doc.json")))

	s.Srv.ListenAndServe()
	return nil
}

func (s *Server) Stop() error {
	s.App.Logg.Info("Stop http server")
	return s.Srv.Close()
}
