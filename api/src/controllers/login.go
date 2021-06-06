package controllers

import (
	"api/src/answers"
	"api/src/auth"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/security"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

func Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		answers.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User

	if err := json.Unmarshal(body, &user); err != nil {
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
	userSaved, err := repo.GetUserByEmail(user.Email)
	if err != nil {
		answers.Error(w, http.StatusInternalServerError, err)
		return
	}

	if err := security.CheckPassword(userSaved.Password, user.Password); err != nil {
		answers.Error(w, http.StatusUnauthorized, err)
		return
	}

	token, err := auth.GetToken(userSaved.ID)
	if err != nil {
		answers.Error(w, http.StatusInternalServerError, err)
		return
	}

	userID := strconv.FormatUint(userSaved.ID, 10)

	answers.JSON(w, http.StatusOK, models.AuthData{ID: userID, Token: token})
}
