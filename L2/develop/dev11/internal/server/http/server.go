package internalhttp

import (
	"errors"
	"fmt"
	"github.com/dingowd/WB/L2/develop/dev11/internal/app"
	"github.com/dingowd/WB/L2/develop/dev11/models"
	"github.com/dingowd/WB/L2/develop/dev11/utils"
	"net/http"
	"strconv"
)

type Server struct {
	//Logg logger.Logger
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

func (s *Server) CreateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(utils.ReturnError("Method isn't POST"))
		return
	}
	var e models.Event
	var err error
	//err = json.NewDecoder(r.Body).Decode(&e)
	err = utils.ToStruct(r, &e)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(utils.ReturnError(err.Error()))
		return
	}
	err = s.App.Storage.Create(e)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(utils.ReturnError(err.Error()))
		return
	}
	w.Write(utils.ReturnResult("Event created"))
}

func (s *Server) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(utils.ReturnError("Method isn't POST"))
		return
	}
	var e models.DBEvent
	var err error
	err = utils.ToDBStruct(r, &e)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		msg := err.Error()
		w.Write(utils.ReturnError(msg))
		return
	}
	err = s.App.Storage.Update(e)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		msg := err.Error()
		w.Write([]byte(msg))
		return
	}
	w.Write(utils.ReturnResult("Event updated"))
}

func (s *Server) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(utils.ReturnError("Method isn`t POST"))
		return
	}
	var e models.DBEvent
	var err error
	err = utils.ToDBStruct(r, &e)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		msg := err.Error()
		w.Write(utils.ReturnError(msg))
		return
	}
	err = s.App.Storage.Delete(e.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		msg := err.Error()
		w.Write(utils.ReturnError(msg))
		return
	}
	w.Write(utils.ReturnResult("Event deleted"))
}

func (s *Server) GetEventsForDay(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(utils.ReturnError("Method isn`t GET"))
		return
	}
	var err error
	var userID int
	userIDS := r.URL.Query().Get("user_id")
	date := r.URL.Query().Get("date")
	userID, err = strconv.Atoi(userIDS)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(utils.ReturnError("Error: Wrong id"))
		return
	}
	var events []models.DBEvent
	events, err = s.App.Storage.GetDayEvent(userID, date)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(utils.ReturnError(err.Error()))
		return
	}
	w.Write(utils.ReturnResultArr(events))
}

func (s *Server) GetEventsForWeek(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(utils.ReturnError("Method isn`t GET"))
		return
	}
	var err error
	var userID int
	userIDS := r.URL.Query().Get("user_id")
	date := r.URL.Query().Get("date")
	userID, err = strconv.Atoi(userIDS)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(utils.ReturnError("Error: Wrong id"))
		return
	}
	var events []models.DBEvent
	events, err = s.App.Storage.GetWeekEvent(userID, date)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(utils.ReturnError(err.Error()))
		return
	}
	w.Write(utils.ReturnResultArr(events))
}

func (s *Server) GetEventsForMonth(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(utils.ReturnError("Method isn`t GET"))
		return
	}
	var err error
	var userID int
	userIDS := r.URL.Query().Get("user_id")
	date := r.URL.Query().Get("date")
	userID, err = strconv.Atoi(userIDS)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(utils.ReturnError("Error: Wrong id"))
		return
	}
	var events []models.DBEvent
	events, err = s.App.Storage.GetMonthEvent(userID, date)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(utils.ReturnError(err.Error()))
		return
	}
	w.Write(utils.ReturnResultArr(events))
}

func (s *Server) Hello(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Write([]byte("Hello from server"))
		return
	}
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprint(w, "Method isn`t GET")
}

func (s *Server) Start() error {
	s.App.Logg.Info("http server starting")
	mux := http.NewServeMux()
	s.Srv = &http.Server{Addr: s.Addr, Handler: mux}
	mux.HandleFunc("/hello", loggingMiddleware(s.Hello, s.App.Logg))
	mux.HandleFunc("/create_event", loggingMiddleware(s.CreateEvent, s.App.Logg))
	mux.HandleFunc("/update_event", loggingMiddleware(s.UpdateEvent, s.App.Logg))
	mux.HandleFunc("/delete_event", loggingMiddleware(s.DeleteEvent, s.App.Logg))
	mux.HandleFunc("/events_for_day", loggingMiddleware(s.GetEventsForDay, s.App.Logg))
	mux.HandleFunc("/events_for_week", loggingMiddleware(s.GetEventsForWeek, s.App.Logg))
	mux.HandleFunc("/events_for_month", loggingMiddleware(s.GetEventsForMonth, s.App.Logg))
	s.Srv.ListenAndServe()
	return nil
}

func (s *Server) Stop() error {
	s.App.Logg.Info("Stop http server")
	return s.Srv.Close()
}
