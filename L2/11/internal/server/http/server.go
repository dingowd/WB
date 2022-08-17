package internalhttp

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/dingowd/otus/hw12_13_14_15_calendar/internal/storage"
)

type Server struct {
	Logg Logger
	App  Application
	Addr string
	Srv  *http.Server
}

type Logger interface {
	Info(msg string)
	Error(msg string)
	Debug(msg string)
	Warn(msg string)
}

type Application interface {
	CreateEvent(ctx context.Context, e storage.Event) error
	UpdateEvent(ctx context.Context, id int, e storage.Event) error
	DeleteEvent(ctx context.Context, id int) error
	GetDayEvent(day string) ([]storage.Event, error)
	GetWeekEvent(day string) ([]storage.Event, error)
	GetMonthEvent(day string) ([]storage.Event, error)
}

func NewServer(logger Logger, app Application, addr string) *Server {
	return &Server{Logg: logger, App: app, Addr: addr}
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

func (s *Server) Hello(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		in, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, err.Error())
			s.Logg.Error(err.Error())
			return
		}
		var req Request
		if err := json.Unmarshal(in, &req); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, err.Error())
			s.Logg.Error(err.Error())
			return
		}
		var res Response
		res.Msg = "Hello, " + req.Msg
		fmt.Fprint(w, res)
		s.Logg.Info(res.Msg)
		return
	}
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprint(w, "Method isn`t GET")
}

func (s *Server) Start(ctx context.Context) error {
	select {
	case <-ctx.Done():
		s.Logg.Error(ErrorStartServer.Error())
		return ErrorStartServer
	default:
		s.Logg.Info("http server starting")
		mux := http.NewServeMux()
		s.Srv = &http.Server{Addr: s.Addr, Handler: mux}
		mux.HandleFunc("/hello", loggingMiddleware(s.Hello, s.Logg))
		if err := s.Srv.ListenAndServe(); err != nil {
			s.Logg.Error(err.Error())
			return err
		}
		return nil
	}
}

func (s *Server) Stop(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ErrorStopServer
	default:
		return s.Srv.Close()
	}
}
