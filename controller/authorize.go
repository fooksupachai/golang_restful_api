package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	database "github.com/fooksupachai/golang_restful_api/database"
	u "github.com/fooksupachai/golang_restful_api/utils"

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
	Type     string `json:"type"`
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

	var jwtKey = []byte(u.GetEnvVariable(`PASSWORD_DATABASE`))

	expirationTime := time.Now().Add(5 * time.Minute)

	claim := &Claims{
		Username: client.Username,
		Type:     "forGetUser",
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

// Auth for check account is exist
func Auth(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	result := database.CheckAccount(r.Body)

	var clients []Client

	for result.Next() {

		var client Client

		_ = result.Scan(
			&client.Username,
			&client.Password,
		)

		clients = append(clients, client)
	}

	if len(clients) != 1 {
		w.WriteHeader(http.StatusForbidden)
		res := u.Message("error")
		res["data"] = "Client doesn't exist"
		u.Response(w, res)
		return
	}

	resp := struct {
		Status int `json:"status"`
	}{
		Status: http.StatusAccepted,
	}

	json.NewEncoder(w).Encode(resp)
}
