package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/src/answers"
	"webapp/src/config"
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

	url := fmt.Sprintf("%s/users", config.APIURL)
	response, err := http.Post(url, "application/json", bytes.NewReader(user))
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
