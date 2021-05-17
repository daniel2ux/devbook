package controllers

import (
	"api/src/answers"
	"api/src/auth"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/security"
	"encoding/json"
	"errors"
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

	userIDToken, err := auth.GetUserID(r)
	if err != nil {
		answers.Error(w, http.StatusUnauthorized, err)
		return
	}

	if userID != userIDToken {
		answers.Error(w, http.StatusForbidden, errors.New("you're only allowed to update your user"))
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

	userIDToken, err := auth.GetUserID(r)
	if err != nil {
		answers.Error(w, http.StatusUnauthorized, err)
		return
	}

	if userID != userIDToken {
		answers.Error(w, http.StatusForbidden, errors.New("you're only allowed to delete your user"))
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

//FollowUser set follow by user
func FollowUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		answers.Error(w, http.StatusBadRequest, err)
		return
	}

	followerID, err := auth.GetUserID(r)
	if err != nil {
		answers.Error(w, http.StatusUnauthorized, err)
		return
	}

	if followerID == userID {
		answers.Error(w, http.StatusForbidden, errors.New("you cannot follow yourself"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		answers.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repositories.UserRepository(db)
	err = repo.Follow(followerID, userID)
	if err != nil {
		answers.Error(w, http.StatusInternalServerError, err)
		return
	}

	answers.JSON(w, http.StatusNoContent, nil)
}

//StopFollowUser stop follow by user
func StopFollowUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		answers.Error(w, http.StatusBadRequest, err)
		return
	}

	followerID, err := auth.GetUserID(r)
	if err != nil {
		answers.Error(w, http.StatusBadRequest, err)
		return
	}

	if followerID == userID {
		answers.Error(w, http.StatusForbidden, errors.New("you cannot stop follow yourself"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		answers.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repositories.UserRepository(db)
	err = repo.StopFollow(followerID, userID)
	if err != nil {
		answers.Error(w, http.StatusInternalServerError, err)
		return
	}

	answers.JSON(w, http.StatusNoContent, nil)
}

//GetFollowers
func GetFollowers(w http.ResponseWriter, r *http.Request) {
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
	followers, err := repo.GetFollowers(userID)
	if err != nil {
		answers.Error(w, http.StatusInternalServerError, err)
		return
	}

	answers.JSON(w, http.StatusOK, followers)
}

//GetFollows
func GetFollows(w http.ResponseWriter, r *http.Request) {
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
	follows, err := repo.GetFollows(userID)
	if err != nil {
		answers.Error(w, http.StatusInternalServerError, err)
		return
	}

	answers.JSON(w, http.StatusOK, follows)
}

//UpdatePass
func UpdatePass(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		answers.Error(w, http.StatusBadRequest, err)
		return
	}

	userIDToken, err := auth.GetUserID(r)
	if err != nil {
		answers.Error(w, http.StatusBadRequest, err)
		return
	}

	if userIDToken != userID {
		answers.Error(w, http.StatusForbidden, errors.New("you cannot update pass"))
		return
	}

	body, err := ioutil.ReadAll(r.Body)

	var pass models.Paswword
	if err := json.Unmarshal(body, &pass); err != nil {
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
	savedPass, err := repo.GetPassByID(userIDToken)
	if err != nil {
		answers.Error(w, http.StatusInternalServerError, err)
		return
	}

	if err := security.CheckPassword(savedPass, pass.Actual); err != nil {
		answers.Error(w, http.StatusUnauthorized, err)
		return
	}

	passWithHash, err := security.Hash(pass.New)
	if err != nil {
		answers.Error(w, http.StatusBadRequest, err)
		return
	}

	if err := repo.UpdatePassword(userID, string(passWithHash)); err != nil {
		answers.Error(w, http.StatusInternalServerError, err)
		return
	}

	answers.JSON(w, http.StatusOK, nil)
}
