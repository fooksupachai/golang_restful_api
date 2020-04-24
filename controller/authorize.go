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
		return
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

	// Provide JWT token when auth success

	var jwtKey = []byte(u.GetEnvVariable(`PASSWORD_DATABASE`))

	expirationTime := time.Now().Add(5 * time.Minute)

	claim := &Claims{
		Username: clients[0].Username,
		Type:     "forGetUser",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	tokenString, err := token.SignedString(jwtKey)

	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

	resp := struct {
		Status int    `json:"status"`
		Token  string `json:"token"`
	}{
		Status: http.StatusAccepted,
		Token:  tokenString,
	}

	json.NewEncoder(w).Encode(resp)
}

// RefreshToken when jwt token expire
func RefreshToken(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {

		w.WriteHeader(http.StatusMethodNotAllowed)
		return

	}

	cookie, err := r.Cookie("token")

	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokenString := cookie.Value

	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return u.GetEnvVariable(`PASSWORD_DATABASE`), nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !token.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// We ensure that a new token is not issued until enough time has elapsed
	// In this case, a new token will only be issued if the old token is within
	// 30 seconds of expiry. Otherwise, return a bad request status
	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Create a new token for the current use, with a renewed expiration time
	expirationTime := time.Now().Add(5 * time.Minute)

	claims.ExpiresAt = expirationTime.Unix()

	tokenNew := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStringNew, err := tokenNew.SignedString(u.GetEnvVariable(`PASSWORD_DATABASE`))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenStringNew,
		Expires: expirationTime,
	})

	resp := struct {
		Status int    `json:"status"`
		Token  string `json:"token"`
	}{
		Status: http.StatusAccepted,
		Token:  tokenStringNew,
	}

	json.NewEncoder(w).Encode(resp)

}
