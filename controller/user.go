package controller

import (
	"encoding/json"
	"net/http"
)

// GetUsers to Get all user in database
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	resp := struct {
		Name string `json:"name"`
	}{
		Name: "fook",
	}

	json.NewEncoder(w).Encode(resp)

}
