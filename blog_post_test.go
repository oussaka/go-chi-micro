package main

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/oussaka/go-chi-micro/handler"
	"github.com/oussaka/go-chi-micro/httphandler"
	"github.com/oussaka/go-chi-micro/model"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

var fakeBlogPosts = []*model.Blogs{{
	ID:              1,
	BlogName:        "Lorem ipsum dolor.",
	BlogDetails:     "Lorem ipsum dolor sit amet consectetur adipiscing elit. Consectetur adipiscing elit quisque faucibus ex sapien vitae. Ex sapien vitae pellentesque sem placerat in id. Placerat in id cursus mi pretium tellus duis. Pretium tellus duis convallis tempus leo eu aenean.",
	BlogDescription: "Lorem ipsum dolor sit amet consectetur adipiscing elit quisque faucibus ex sapien vitae pellentesque sem.",
}}

var BlogData = []model.BlogData{{
	Blog:    *fakeBlogPosts[0],
	Message: "OK",
}}

var fakeResponse = httphandler.Response{
	Data: *fakeBlogPosts[0],
}

type fakeStorage struct {
}

func (s fakeStorage) Get(_ string) (model.Blogs, error) {
	return *fakeBlogPosts[0], nil
}

func (s fakeStorage) Delete(_ string) (string, error) {
	return "OK deleted", nil
}

func (s fakeStorage) List() []*model.Blogs {
	return fakeBlogPosts
}

func (s fakeStorage) Create(_ *model.Blogs) (model.BlogData, error) {
	return BlogData[0], nil
}

func (s fakeStorage) Update(_ string, _ *model.Blogs) (model.BlogData, error) {
	return BlogData[0], nil
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

	//Post := model.Post{}
	//json.Unmarshal(data, &Post)

	json.Unmarshal(data, &fakeResponse)

	var returnedBlogName = fakeResponse.Data.(map[string]interface{})["blog_name"]
	if returnedBlogName != "Lorem ipsum dolor." {
		t.Errorf("expected 'Lorem ipsum dolor.' got %v", returnedBlogName)
	}
	if fakeResponse.Data.(map[string]interface{})["id"] != float64(1) {
		t.Errorf("expected Post ID = 1 got %v", fakeResponse.Data.(map[string]interface{})["id"])
	}
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
