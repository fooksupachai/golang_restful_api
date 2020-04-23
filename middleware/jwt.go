package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	u "github.com/fooksupachai/golang_restful_api/utils"
)

// JWTMiddleware a jwt middleware that verify JWT
func JWTMiddleware(handleFunc func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Add("Content-Type", "application/json")

		var header = r.Header.Get("Authorization")

		header = strings.TrimSpace(header)

		headerSlice := strings.Split(header, " ")

		if len(headerSlice) < 2 {

			w.WriteHeader(http.StatusForbidden)

			res := u.Message("error")

			res["data"] = "Missing auth token"

			u.Response(w, res)

			return
		}

		tokenString := headerSlice[1]

		// Parse takes the token string and a function for looking up the key. The latter is especially
		// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
		// head of the token to identify which key to use, but the parsed token (head and claims) is provided
		// to the callback, providing flexibility.
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(u.GetEnvVariable(`PASSWORD_DATABASE`)), nil

		})

		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			res := u.Message("error")
			res["data"] = err.Error()
			u.Response(w, res)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Check that token is what we're expecting
			if claims["type"] == u.GetEnvVariable(`JWT_USER`) {
				handleFunc(w, r)
			} else {
				w.WriteHeader(http.StatusForbidden)
				res := u.Message("error")
				res["data"] = "Unknow client"
				u.Response(w, res)
				return
			}
		} else {
			w.WriteHeader(http.StatusForbidden)
			res := u.Message("error")
			res["data"] = "Invalid token"
			u.Response(w, res)
			return
		}
	}
}
