package utils

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// Stand alone writer interface to from web socket
var write http.ResponseWriter

// Message package a message
func Message(message string) map[string]interface{} {
	return map[string]interface{}{"message": message}
}

// Response create a HTTP response.
func Response(w http.ResponseWriter, data map[string]interface{}) {

	w.Header().Add("Content-Type", "application/json")

	json.NewEncoder(w).Encode(data)
}

// GetEnvVariable from .env file
func GetEnvVariable(key string) string {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv(key)
}

// ValiateQueryURL from client
func ValiateQueryURL(w http.ResponseWriter, param string) {
	res := Message("error")
	res["data"] = "Missing " + param + " parameter"
	Response(w, res)
	return
}

// SocketMessage to handle web socket activity
func SocketMessage(message error) {

	resp := struct {
		Message error `json:"message"`
		Status  int   `json:"status"`
	}{
		Message: message,
		Status:  http.StatusInternalServerError,
	}

	json.NewEncoder(write).Encode(resp)

}
