package internalhttp

import (
	"context"
	"errors"
	"fmt"
	"github.com/dingowd/WB/L2/develop/dev11/internal/app"
	"net/http"
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
		w.Write([]byte("Method isn`t POST"))
		return
	}
}

func (s *Server) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Write([]byte("Method isn`t POST"))
		return
	}
}

func (s *Server) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Write([]byte("Method isn`t POST"))
		return
	}
}

func (s *Server) GetEventsForDay(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) GetEventsForWeek(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) GetEventsForMonth(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) Hello(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Write([]byte("Hello from server"))
		return
	}
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprint(w, "Method isn`t GET")
}

func (s *Server) Start(ctx context.Context) error {
	select {
	case <-ctx.Done():
		s.App.Logg.Error(ErrorStartServer.Error())
		return ErrorStartServer
	default:
		s.App.Logg.Info("http server starting")
		mux := http.NewServeMux()
		s.Srv = &http.Server{Addr: s.Addr, Handler: mux}
		mux.HandleFunc("/hello", s.Hello)
		if err := s.Srv.ListenAndServe(); err != nil {
			s.App.Logg.Error(err.Error())
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
