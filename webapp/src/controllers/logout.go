package controllers

import (
	"net/http"
	"webapp/src/cookies"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	cookies.Remove(w)
	http.Redirect(w, r, "/login", 302)
}
