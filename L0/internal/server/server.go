package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dingowd/WB/L0/model"
	"github.com/dingowd/WB/L0/provider"
	"io/ioutil"
	"net/http"
)

type Server struct {
	Logg     model.Logger
	App      Application
	Addr     string
	Srv      *http.Server
	Provider provider.MessageProvider
}

var (
	ErrorStopServer  = errors.New("timeout to stop server")
	ErrorStartServer = errors.New("timeout to start server")
)

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
