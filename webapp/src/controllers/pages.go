package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"webapp/src/answers"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/models"
	"webapp/src/requests"
	"webapp/src/utils"

	"github.com/gorilla/mux"
)

func LoadLoginScreen(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Read(r)

	if cookie["token"] != "" {
		http.Redirect(w, r, "/home", 302)
		return
	}

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

	cookie, _ := cookies.Read(r)
	userID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	utils.LoadTemplate(w, "home.html", struct {
		Posts  []models.Post
		UserID uint64
	}{
		Posts:  posts,
		UserID: userID,
	})
}

func LoadEditPage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	postID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		answers.JSON(w, http.StatusBadRequest, answers.APIError{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/posts/%d", config.APIURL, postID)
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

	var post models.Post
	if err := json.NewDecoder(response.Body).Decode(&post); err != nil {
		answers.JSON(w, http.StatusUnprocessableEntity, answers.APIError{Error: err.Error()})
		return
	}

	utils.LoadTemplate(w, "edit-post.html", post)
}
