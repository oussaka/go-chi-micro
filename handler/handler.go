package handler

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/oussaka/go-chi-micro/model"
	"net/http"
)

type BlogHandler struct {
	Storage BlogStorage
}

type ctx struct {
	storage BlogStorage
	h       func(BlogStorage, http.ResponseWriter, *http.Request)
}

func (g *ctx) handle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		g.h(g.storage, w, r)
	}
}

func Handler(handler *BlogHandler) http.Handler {
	r := chi.NewRouter()
	getRecordSetPost := ctx{storage: handler.Storage, h: handler.GetPosts}
	createBlogPost := ctx{storage: handler.Storage, h: handler.CreatePost}
	updateBlogs := ctx{storage: handler.Storage, h: handler.UpdatePost}
	deleteBlogs := ctx{storage: handler.Storage, h: handler.DeletePost}

	r.Get(wrapHandlerFunc("/blog/{id}", "get Post", getRecordSetPost.handle()))
	r.Post(wrapHandlerFunc("/blog", "create blog", createBlogPost.handle()))
	r.Put(wrapHandlerFunc("/blog/{data}", "update blog", updateBlogs.handle()))
	r.Delete(wrapHandlerFunc("/blog/remove/{data}", "delete blog", deleteBlogs.handle()))

	return r
}

func (b BlogStore) ListPosts(service BlogStore, w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(b.List())
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}

func (b BlogHandler) GetPosts(service BlogStorage, w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	Post := b.Storage.Get(id)
	if Post == nil {
		http.Error(w, "Post not found", http.StatusNotFound)
	}
	err := json.NewEncoder(w).Encode(Post)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}

func (b BlogHandler) CreatePost(service BlogStorage, w http.ResponseWriter, r *http.Request) {
	var Post model.Post
	err := json.NewDecoder(r.Body).Decode(&Post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	b.Storage.Create(Post)
	err = json.NewEncoder(w).Encode(Post)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}

func (b BlogHandler) UpdatePost(service BlogStorage, w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var Post model.Post
	err := json.NewDecoder(r.Body).Decode(&Post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	updatedPost := b.Storage.Update(id, Post)
	if updatedPost == nil {
		http.Error(w, "Post not found", http.StatusNotFound)
	}
	err = json.NewEncoder(w).Encode(updatedPost)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}

func (b BlogHandler) DeletePost(service BlogStorage, w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	Post := b.Storage.Delete(id)
	if Post == nil {
		http.Error(w, "Post not found", http.StatusNotFound)
	}
	w.WriteHeader(http.StatusNoContent)
}

func wrapHandlerFunc(route string, name string, handler http.HandlerFunc) (string, http.HandlerFunc) {
	return route, handler
}
