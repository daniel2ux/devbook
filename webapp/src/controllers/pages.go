package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/src/answers"
	"webapp/src/config"
	"webapp/src/models"
	"webapp/src/requests"
	"webapp/src/utils"
)

func LoadLoginScreen(w http.ResponseWriter, r *http.Request) {
	utils.LoadTemplate(w, "login.html", nil)
}

func LoadUserEnrollPage(w http.ResponseWriter, r *http.Request) {
	utils.LoadTemplate(w, "enroll.html", nil)
}

func LoadMainPage(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/posts", config.APIURL)
	response, err := requests.RequestWithAuth(r, http.MethodGet, url, nil)
	if err != nil {
		answers.JSON(w, http.StatusInternalServerError, answers.APIError{Error: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		answers.CheckStatusCodeError(w, response)
		return
	}

	var posts []models.Post

	if err := json.NewDecoder(response.Body).Decode(&posts); err != nil {
		answers.JSON(w, http.StatusBadRequest, err)
		return
	}

	utils.LoadTemplate(w, "home.html", posts)
}
