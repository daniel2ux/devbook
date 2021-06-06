package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/src/answers"
	"webapp/src/config"
	"webapp/src/models"
)

func DoLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	user, err := json.Marshal(map[string]string{
		"email":    r.FormValue("email"),
		"password": r.FormValue("password"),
	})
	if err != nil {
		answers.JSON(w, http.StatusBadRequest, answers.APIError{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/login", config.APIURL)
	response, err := http.Post(url, "application/json", bytes.NewBuffer(user))
	if err != nil {
		answers.JSON(w, http.StatusInternalServerError, answers.APIError{Error: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		answers.CheckStatusCodeError(w, response)
		return
	}

	var authData models.AuthData
	if err := json.NewDecoder(response.Body).Decode(&authData); err != nil {
		answers.JSON(w, http.StatusUnprocessableEntity, answers.APIError{Error: err.Error()})
		return
	}

	answers.JSON(w, http.StatusOK, nil)
}
