package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	database "github.com/fooksupachai/golang_restful_api/database"

	"github.com/dgrijalva/jwt-go"
)

// Client for user handler
type Client struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Claims for jwt client
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// CreateAccount into app
func CreateAccount(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")

	if r.Method != http.MethodPost {

		w.WriteHeader(http.StatusMethodNotAllowed)

		return
	}

	var client Client

	json.NewDecoder(r.Body).Decode(&client)

	if client.Password == "" || client.Username == "" {

		respNotify := struct {
			Notify string `json:"notify"`
			Status int    `json:"status"`
		}{
			Notify: "Username or Password cannot emthy",
			Status: http.StatusBadRequest,
		}

		json.NewEncoder(w).Encode(respNotify)

		return
	}

	var jwtKey = []byte("fyhwek+1t(aE")

	expirationTime := time.Now().Add(1 * time.Minute)

	claim := &Claims{
		Username: client.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	tokenString, err := token.SignedString(jwtKey)

	fmt.Println(err)

	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	database.CreateAccontData(
		client.Username,
		client.Password,
	)

	fmt.Println(r.Body)

	respAccount := struct {
		Username string `json:"username"`
		Token    string `json:"token"`
	}{
		Username: client.Username,
		Token:    tokenString,
	}

	json.NewEncoder(w).Encode(respAccount)

}
