package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type ShutdownFn func() error

type Server struct {
	srv        *http.Server
	shutdownFn ShutdownFn
}

func NewServer(addr string, r *mux.Router) *Server {
	server := &http.Server{
		Addr:         addr,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprint(w, "OK")
		if err != nil {
			return
		}
	})

	return &Server{
		srv: server,
	}
}

func (s *Server) ListenAndServeForever() {
	s.shutdownFn = func() error {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
		defer cancel()
		return s.srv.Shutdown(ctx)
	}

	if err := s.srv.ListenAndServe(); err != nil &&
		!errors.Is(err, http.ErrServerClosed) {
		log.Println(err)
	}
}

func (s *Server) WaitForShutdown() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	s.Shutdown()
}

func (s *Server) Shutdown() {
	if err := s.shutdownFn(); err != nil {
		log.Println(err)
		os.Exit(1)
	}

	log.Println("App shut down gracefully")
	os.Exit(0)
}

func (s *Server) Address() string {
	return s.srv.Addr
}
