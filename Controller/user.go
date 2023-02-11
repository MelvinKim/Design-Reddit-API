package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	model "github.com/MelvinKim/Design-Reddit-API/Model"
	usecase "github.com/MelvinKim/Design-Reddit-API/Usecase"
)

type userController struct{}

var (
	userUsecase usecase.UserUsecase
)

func NewUserController(usecase usecase.UserUsecase) UserController {
	userUsecase = usecase
	return &userController{}
}

func (c *userController) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorResponse := map[string]string{"error": "Bad request"}
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	result, err1 := userUsecase.Create(&user)
	if err1 != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorResponse := map[string]string{"error": "error saving the post"}
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}

func (c *userController) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Path[len("/users/"):]
	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	user, err1 := userUsecase.Get(idInt)
	if err1 != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorResponse := map[string]string{"error": fmt.Sprintf("error getting user with id: %v", idInt)}
		json.NewEncoder(w).Encode(errorResponse)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (c *userController) FindAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	users, err := userUsecase.FindAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorResponse := map[string]string{"error": "error fetching users"}
		json.NewEncoder(w).Encode(errorResponse)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func (c *userController) Update(w http.ResponseWriter, r *http.Request) {

}

func (c *userController) Delete(w http.ResponseWriter, r *http.Request) {}
