package main

import (
	"github.com/oussaka/go-chi-micro/api"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type Config struct{}

func main() {
	app := Config{}
	r := app.routes()
	s := server.New()

	//log.Info("Listening on port:", config.GetYamlValues().ServerConfig.Port)
	err := s.ListenAndServe()
	if err != http.ErrServerClosed {
		log.Fatalf("Listen: %s\n", err)
	}

	log.Info("service stopped")
}
