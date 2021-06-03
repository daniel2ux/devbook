package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"webapp/src/answers"
)

func EnrollUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	user, err := json.Marshal(map[string]string{
		"name":     r.FormValue("name"),
		"email":    r.FormValue("email"),
		"nick":     r.FormValue("nick"),
		"password": r.FormValue("password"),
	})

	if err != nil {
		answers.JSON(w, http.StatusBadRequest, answers.APIError{Error: err.Error()})
		return
	}

	response, err := http.Post("http://localhost:5000/users", "application/json", bytes.NewReader(user))
	if err != nil {
		answers.JSON(w, http.StatusInternalServerError, answers.APIError{Error: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		answers.CheckStatusCodeError(w, response)
		return
	}

	answers.JSON(w, response.StatusCode, nil)
}
