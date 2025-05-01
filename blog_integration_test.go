package main

import (
	"fmt"
	"github.com/oussaka/go-chi-micro/server"
	"net/http"
	"net/http/httptest"
	"testing"
)

func runTestServer() *httptest.Server {
	return httptest.NewServer(server.New().GetHandler())
}

func TestIntegrationGetBlogPostsHandler(t *testing.T) {
	testServer := runTestServer()
	defer testServer.Close()

	resp, err := http.Get(fmt.Sprintf("%s/api/v1/blog/1", testServer.URL))

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != 200 {
		t.Errorf("expecte 200 got: %v", resp.StatusCode)
	}
}
