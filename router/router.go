package router

import (
	"log"
	"net/http"

	controller "github.com/fooksupachai/golang_restful_api/controller"
)

func init() {

	mux := http.NewServeMux()

	mux.HandleFunc("/users", controller.GetUsers)
	mux.HandleFunc("/user", controller.CreateUser)
	mux.HandleFunc("/user/", controller.UpdateUser) // to handle user/:firstname.

	log.Print("Rest server is listening on port 8080")
	http.ListenAndServe(":8080", mux)
}
