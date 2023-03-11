package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestHomePageHandler(t *testing.T) {
	mockResponse := `{"message":"Welcome to the Reddit API build with Golang"}`
	r := SetUpRouter()
	r.GET("/", HomepageHandler)
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Error("error while creating request for HomepageHandler: ", err)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, err := io.ReadAll(w.Body)
	if err != nil {
		t.Error("error while reading HomepageHandler response body: ", err)
	}
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}
