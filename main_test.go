package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/MelvinKim/Design-Reddit-API/entity"
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

	responseData, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Error("error while reading HomepageHandler response body: ", err)
	}
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCreateUserHandler(t *testing.T) {
	r := SetUpRouter()
	r.POST("/users", UserController.CreateUser)
	userID := 49
	user := &entity.User{
		ID:        uint32(userID),
		FirstName: "test",
		LastName:  "test",
		Email:     "test@test.com",
		Password:  "test",
		IsDeleted: false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	jsonValue, err := json.Marshal(user)
	if err != nil {
		t.Error("error while marshalling user struct to JSON: ", err)
		return
	}
	req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonValue))
	if err != nil {
		t.Error("error while creating CreateUserHandler request: ", err)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}
