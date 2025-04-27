package main

import (
	"github.com/go-chi/chi/v5/middleware"
	"github.com/oussaka/go-chi-micro/handler"
	"github.com/oussaka/go-chi-micro/model"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := setupServer()
	http.ListenAndServe(":3000", r)
}

func setupServer() chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Heartbeat("/ping"))
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	r.Mount("/v1", BlogRoutes())

	return r
}

func BlogRoutes() chi.Router {
	r := chi.NewRouter()

	blogHandler := handler.BlogHandler{
		Storage: model.BlogStore{},
	}

	r.Get("/", blogHandler.ListPosts)
	r.Post("/", blogHandler.CreatePost)
	r.Get("/{id}", blogHandler.GetPosts)
	r.Put("/{id}", blogHandler.UpdatePost)
	r.Delete("/{id}", blogHandler.DeletePost)

	return r
}
