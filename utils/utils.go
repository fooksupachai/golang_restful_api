package utils

import (
	"encoding/json"
	"net/http"
)

// Message package a message
func Message( message string) map[string]interface{}{
	return map[string]interface{}{"message": message}
}

// Response create a HTTP response.
func Response(w http.ResponseWriter, data map[string]interface{}) {
	
	w.Header().Add("Content-Type", "application/json")

	json.NewEncoder(w).Encode(data)
}