package answers

import (
	"encoding/json"
	"log"
	"net/http"
)

type APIError struct {
	Error string `json:"error"`
}

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal(err)
	}
}

func CheckStatusCodeError(w http.ResponseWriter, r *http.Response) {
	var error APIError
	json.NewDecoder(r.Body).Decode(&error)
	JSON(w, r.StatusCode, error)
}
