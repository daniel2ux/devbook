package controllers

import (
	"net/http"
	"webapp/src/utils"
)

func LoadLoginScreen(w http.ResponseWriter, r *http.Request) {
	utils.LoadTemplate(w, "login.html", nil)
}
