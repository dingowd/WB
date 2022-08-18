package internalhttp

import (
	"encoding/json"
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
	err = json.NewDecoder(r.Body).Decode(&e)
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
	/*	idS := r.URL.Query().Get("id")
		if len(idS) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(utils.ReturnError("id is missing. check request."))
			return
		}
		var err error
		var id int
		id, err = strconv.Atoi(idS)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(utils.ReturnError("Wrong id"))
			return
		}*/
	var e models.DBEvent
	var err error
	err = json.NewDecoder(r.Body).Decode(&e)
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
	idS := r.URL.Query().Get("id")
	if len(idS) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(utils.ReturnError("id is missing. check request."))
		return
	}
	var err error
	var id int
	id, err = strconv.Atoi(idS)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(utils.ReturnError("Wrong id"))
		return
	}
	err = s.App.Storage.Delete(id)
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
	mux.HandleFunc("/hello", s.Hello)
	mux.HandleFunc("/create_event", s.CreateEvent)
	mux.HandleFunc("/update_event", s.UpdateEvent)
	mux.HandleFunc("/delete_event", s.DeleteEvent)
	mux.HandleFunc("/events_for_day", s.GetEventsForDay)
	mux.HandleFunc("/events_for_week", s.GetEventsForWeek)
	mux.HandleFunc("/events_for_month", s.GetEventsForMonth)
	if err := s.Srv.ListenAndServe(); err != nil {
		s.App.Logg.Error(err.Error())
		return err
	}
	return nil
}

func (s *Server) Stop() error {
	s.App.Logg.Info("Stop http server")
	return s.Srv.Close()
}
