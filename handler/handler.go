package handler

import (
	"encoding/json"
	"errors"
	"github.com/jinzhu/gorm"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/oussaka/go-chi-micro/httphandler"
	"github.com/oussaka/go-chi-micro/model"
	log "github.com/sirupsen/logrus"
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
	r.Use(render.SetContentType(render.ContentTypeJSON))

	getRecordSetPost := ctx{storage: handler.Storage, h: handler.GetPosts}
	createBlogPost := ctx{storage: handler.Storage, h: handler.CreatePost}
	updateBlogs := ctx{storage: handler.Storage, h: handler.UpdatePost}
	deleteBlogs := ctx{storage: handler.Storage, h: handler.DeletePost}

	r.Get(httphandler.WrapHandlerFunc("/blog/{id}", "get Post", getRecordSetPost.handle()))
	r.Post(httphandler.WrapHandlerFunc("/blog", "create blog", createBlogPost.handle()))
	r.Put(httphandler.WrapHandlerFunc("/blog/{id}", "update blog", updateBlogs.handle()))
	r.Delete(httphandler.WrapHandlerFunc("/blog/remove/{id}", "delete blog", deleteBlogs.handle()))

	return r
}

func (b BlogHandler) ListPosts(service BlogStorage, w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(b.Storage.List())
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}

func (b BlogHandler) GetPosts(service BlogStorage, w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	post, err := b.Storage.Get(id)
	if err != nil {
		log.Errorf("Unable To Fetch stats ", httphandler.Error(err).Code, post, err)

		if errors.Is(err.(error), gorm.ErrRecordNotFound) {
			render.Render(w, r, httphandler.ErrNotFoundRequest(err, err.Error()))
		} else {
			render.Render(w, r, httphandler.ErrInvalidRequest(err, "Unable To Fetch Services "))
		}
		return
	}

	render.Status(r, http.StatusOK)
	render.Render(w, r, httphandler.NewSuccessResponse(http.StatusOK, post))
}

func (b BlogHandler) CreatePost(service BlogStorage, w http.ResponseWriter, r *http.Request) {
	var Post *model.Blogs
	err := json.NewDecoder(r.Body).Decode(&Post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	postData, err := b.Storage.Create(Post)
	if err != nil {
		log.Errorf("Unable To Fetch stats ", httphandler.Error(err).Code, postData, err)
		httphandler.ErrInvalidRequest(err, "Unable To Fetch Services ")
		return
	}

	render.Status(r, http.StatusOK)
	render.Render(w, r, httphandler.NewSuccessResponse(http.StatusOK, postData))
}

func (b BlogHandler) UpdatePost(service BlogStorage, w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var Post model.Blogs
	err := json.NewDecoder(r.Body).Decode(&Post)
	if err != nil {
		render.Render(w, r, httphandler.ErrInvalidRequest(err, err.Error()))
		return
	}
	updatedPost, err := b.Storage.Update(id, &Post)

	if err != nil {
		log.Errorf("Unable To Fetch stats ", httphandler.Error(err).Code, updatedPost, err)
		httphandler.ErrInvalidRequest(err, "Unable To Fetch Services ")
		return
	}

	render.Status(r, http.StatusOK)
	render.Render(w, r, httphandler.NewSuccessResponse(http.StatusOK, updatedPost))
}

func (b BlogHandler) DeletePost(service BlogStorage, w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	post, err := b.Storage.Delete(id)
	if err != nil {
		log.Errorf("Unable To Fetch stats ", httphandler.Error(err).Code, post, err)
		httphandler.ErrInvalidRequest(err, "Unable To Fetch Services ")
		return
	}

	render.Status(r, http.StatusOK)
	render.Render(w, r, httphandler.NewSuccessResponse(http.StatusOK, post))
}
