package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//CreateUser create a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var user models.User

	if err = json.Unmarshal(body, &user); err != nil {
		log.Fatal(err)
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	repository := repositories.UserRepository(db)
	userID, err := repository.Create(user)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("User created... ID: %d", userID)))
}

//GetUsers get all users from database
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting users..."))
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
