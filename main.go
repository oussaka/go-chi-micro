package main

import (
	"github.com/oussaka/go-chi-micro/server"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	s := server.New()
	err := s.ListenAndServe()
	if err != http.ErrServerClosed {
		log.Fatalf("Listen: %s\n", err)
	}

	log.Info("service stopped")
}
