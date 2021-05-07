package controllers

import (
	"api/src/answers"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
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

	if err := user.Prepare("add"); err != nil {
		answers.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		answers.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repositories.UserRepository(db)
	user.ID, err = repo.Create(user)
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

	repo := repositories.UserRepository(db)
	users, err := repo.GetUsers(nameOrNick)
	if err != nil {
		answers.Error(w, http.StatusInternalServerError, err)
		return
	}

	answers.JSON(w, http.StatusOK, users)
}

//GetUser gets specific user from database
func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		answers.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		answers.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repositories.UserRepository(db)
	user, err := repo.GetUserByID(userID)
	if err != nil {
		answers.Error(w, http.StatusInternalServerError, err)
		return
	}

	answers.JSON(w, http.StatusOK, user)
}

//UpdateUser update data from specific user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		answers.Error(w, http.StatusBadRequest, err)
		return
	}

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

	if err := user.Prepare("edit"); err != nil {
		answers.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		answers.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repositories.UserRepository(db)
	err = repo.Update(userID, user)
	if err != nil {
		answers.Error(w, http.StatusInternalServerError, err)
		return
	}

	answers.JSON(w, http.StatusNoContent, nil)

}

//DeleteUser delete specific user from database
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		answers.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		answers.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repositories.UserRepository(db)
	err = repo.Delete(userID)
	if err != nil {
		answers.Error(w, http.StatusInternalServerError, err)
		return
	}

	answers.JSON(w, http.StatusNoContent, nil)
}
