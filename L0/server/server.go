package server

import (
	"fmt"
	"github.com/dingowd/WB/L0/app"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	App  *app.App
	Addr string
	Srv  *http.Server
}

func NewServer(app *app.App, addr string) *Server {
	s := &Server{App: app, Addr: addr}
	return s
}

func (s *Server) GetOrder(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if len(id) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("id is missing!!!"))
		return
	}
	order, err := s.App.Cache.ReadFromCache(id)
	if err != nil {
		msg := err.Error() + " " + id
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(msg))
		s.App.Log.Error(msg)
		return
	}
	tmpl, errT := template.ParseFiles("./L0/templates/index.html")
	if errT != nil {
		s := fmt.Sprint(order)
		w.Write([]byte(s))
	} else {
		tmpl.Execute(w, order.Order)
	}
}

func (s *Server) Start() error {
	router := mux.NewRouter()
	router.HandleFunc("/get", s.GetOrder).Methods("GET")
	http.Handle("/", router)
	Srv := &http.Server{Addr: s.Addr, Handler: router}
	s.Srv = Srv
	s.App.Log.Info("http сервер запускается")
	err := s.Srv.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) Stop() error {
	s.App.Log.Info("остановка http сервера")
	return s.Srv.Close()
}
