package main

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/oussaka/go-chi-micro/handler"
	"github.com/oussaka/go-chi-micro/model"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

var fakePosts = []*model.Post{{
	ID:               "1",
	Title:            "7 habits of highly effective people",
	Author:           "Stephen Covey",
	PublishedDate:    "15/08/1989",
	OriginalLanguage: "english",
}}

type fakeStorage struct {
}

func (s fakeStorage) Get(_ string) *model.Post {
	return fakePosts[0]
}

func (s fakeStorage) Delete(_ string) *model.Post {
	return fakePosts[0]
}

func (s fakeStorage) List() []*model.Post {
	return fakePosts
}

func (s fakeStorage) Create(_ model.Post) {
	return
}

func (s fakeStorage) Update(_ string, _ model.Post) *model.Post {
	return fakePosts[1]
}

func TestGetPostsHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/v1/blog/1", nil)
	w := httptest.NewRecorder()
	blogHandler := &handler.BlogHandler{
		Storage: fakeStorage{},
	}

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "1")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	blogHandler.GetPosts(blogHandler.Storage, w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}

	// Check the response code
	checkResponseCode(t, http.StatusOK, res.StatusCode)

	Post := model.Post{}
	json.Unmarshal(data, &Post)
	if Post.Title != "7 habits of highly effective people" {
		t.Errorf("expected ABC got %v", string(data))
	}
	if Post.ID != "1" {
		t.Errorf("expected Post ID = 1 got %v", Post.ID)
	}
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
