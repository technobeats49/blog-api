package controllers

import (
	"encoding/json"
	"example.com/blog-api/db"
	"example.com/blog-api/models"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAllArticlesJSON(t *testing.T) {
	resp, err := http.Get("http://localhost:9000/articles")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 200 {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
}

func TestGetAllArticles(t *testing.T) {
	r := httprouter.New()
	bsc := NewBlogStoreController(db.DBS)
	r.GET("/articles", bsc.GetAllArticles)
	req, _ := http.NewRequest("GET", "/articles", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	art := []models.Article{}
	json.Unmarshal(w.Body.Bytes(), &art)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, art)
}
