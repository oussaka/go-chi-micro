package handler

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/oussaka/go-chi-micro/model"
	"net/http"
)

type BlogHandler struct {
	Storage model.BlogStorage
}

func (b BlogHandler) ListPosts(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(b.Storage.List())
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}

func (b BlogHandler) GetPosts(w http.ResponseWriter, r *http.Request) {
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

func (b BlogHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
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

func (b BlogHandler) UpdatePost(w http.ResponseWriter, r *http.Request) {
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

func (b BlogHandler) DeletePost(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	Post := b.Storage.Delete(id)
	if Post == nil {
		http.Error(w, "Post not found", http.StatusNotFound)
	}
	w.WriteHeader(http.StatusNoContent)
}
