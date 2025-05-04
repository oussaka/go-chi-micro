package server

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/oussaka/go-chi-micro/db"
	"github.com/oussaka/go-chi-micro/handler"
	"net"
	"net/http"
	"os"
)

const webPort = "3000"

type Server struct {
	httpServer *http.Server
	//router     *chi.Mux
}

func New() *Server {
	r := chi.NewRouter()

	db.InitPgsql()

	sqlClient := db.NewClient(
		&db.Config{
			DBConnection: "",
		})

	//sqlConn, err := db.CreateConnection()
	//if err != nil {
	//	log.Panic("Can't connect to Postgres !", err)
	//}

	//dbConfig := &db.Config{
	//	DbPool: sqlConn.DbPool,
	//}
	//sqlClient := db.NewClient(dbConfig)
	storage := handler.NewService(sqlClient)

	blogHandler := &handler.BlogHandler{Storage: storage}
	setupRoutes(blogHandler, r)

	server := newServer(r)

	return server
}

func setupRoutes(service *handler.BlogHandler, r *chi.Mux) {
	// plug in sub-routers for resources: feature gate
	// this pattern also allows for easy integration testing. see api_test.go

	r.Route("/api", func(r chi.Router) {
		r.Mount("/v1", handler.Handler(service))
	})
}

func (s *Server) ListenAndServe() error {
	l, err := net.Listen("tcp", ":"+s.httpServer.Addr)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}

	return s.httpServer.Serve(l)
}

func newServer(r http.Handler) *Server {
	fmt.Println("****Server Started on", webPort, "****")
	return &Server{
		httpServer: &http.Server{
			Addr:    fmt.Sprintf("%s", webPort),
			Handler: r,
		},
	}
}

func (s *Server) GetHandler() http.Handler {
	return s.httpServer.Handler
}
