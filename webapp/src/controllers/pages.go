package controllers

import (
	"net/http"
	"webapp/src/utils"
)

func LoadLoginScreen(w http.ResponseWriter, r *http.Request) {
	utils.LoadTemplate(w, "login.html", nil)
}

func LoadUserEnrollPage(w http.ResponseWriter, r *http.Request) {
	utils.LoadTemplate(w, "enroll.html", nil)
}
