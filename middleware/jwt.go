package middleware 

import (
	"net/http"
	"strings"

	u "github.com/fooksupachai/golang_restful_api/utils"
)

// JWTMiddleware a jwt middleware that verify JWT
func JWTMiddleware(handleFunc func(http.ResponseWriter, *http.Request)) http.HandlerFunc{
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

		

	}
}