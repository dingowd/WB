package server

import (
	"encoding/json"
	"errors"
	"github.com/dingowd/WB/L0/app"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	App  *app.App
	Addr string
	Srv  *http.Server
}

var (
	ErrorStopServer  = errors.New("timeout to stop server")
	ErrorStartServer = errors.New("timeout to start server")
)

func NewServer(app *app.App, addr string) *Server {
	s := &Server{App: app, Addr: addr}
	return s
}

func (s *Server) GetOrder(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if len(id) == 0 {
		w.Write([]byte("id is missing!!!"))
		return
	}
	order, err := s.App.Cache.ReadFromCache(id)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	b, _ := json.Marshal(order)
	w.Write(b)
}

func (s *Server) Start() error {
	router := mux.NewRouter()
	router.HandleFunc("/get", s.GetOrder).Methods("GET")
	http.Handle("/", router)
	Srv := &http.Server{Addr: s.Addr, Handler: router}
	s.App.Log.Info("http сервер запускается")
	err := Srv.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) Stop() error {
	return s.Srv.Close()
}
