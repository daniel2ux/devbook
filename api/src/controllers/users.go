package controllers

import (
	"api/src/answers"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

//CreateUser create a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		answers.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(body, &user); err != nil {
		answers.Error(w, http.StatusBadRequest, err)
		return
	}

	if err := user.Prepare(); err != nil {
		answers.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		answers.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.UserRepository(db)
	user.ID, err = repository.Create(user)
	if err != nil {
		answers.Error(w, http.StatusInternalServerError, err)
		return
	}

	answers.JSON(w, http.StatusCreated, user)
}

//GetUsers get all users from database
func GetUsers(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))

	db, err := database.Connect()
	if err != nil {
		answers.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.UserRepository(db)
	users, err := repository.GetUsers(nameOrNick)
	if err != nil {
		answers.Error(w, http.StatusInternalServerError, err)
		return
	}

	answers.JSON(w, http.StatusOK, users)
}

//GetUser gets specific user from database
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting user..."))
}

//UpdateUser update data from specific user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update user..."))
}

//DeleteUser delete specific user from database
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete user..."))
}
