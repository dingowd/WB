package server

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/dingowd/WB/L0/app"
	"net/http"
)

type Server struct {
	App  app.App
	Addr string
	Srv  *http.Server
}

var (
	ErrorStopServer  = errors.New("timeout to stop server")
	ErrorStartServer = errors.New("timeout to start server")
)

func NewServer(app app.App, addr string) *Server {
	return &Server{App: app, Addr: addr}
}

func (s *Server) GetOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.Write([]byte("Method is not GET!!!"))
		return
	}
	id := r.URL.Query().Get("id")
	if len(id) == 0 {
		w.Write([]byte("id is missing!!!"))
		return
	}
	order := s.App.Cache.ReadFromCache(id)
	b, _ := json.Marshal(order)
	w.Write(b)
}

func (s *Server) Start(ctx context.Context) error {
	select {
	case <-ctx.Done():
		s.App.Log.Error(ErrorStartServer.Error())
		return ErrorStartServer
	default:
		s.App.Log.Info("http server starting")
		mux := http.NewServeMux()
		s.Srv = &http.Server{Addr: s.Addr, Handler: mux}
		mux.HandleFunc("/get", s.GetOrder)
		if err := s.Srv.ListenAndServe(); err != nil {
			s.App.Log.Error(err.Error())
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
