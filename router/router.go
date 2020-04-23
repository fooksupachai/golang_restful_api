package router

import (
	"net/http"

	"github.com/gorilla/mux"

	controller "github.com/fooksupachai/golang_restful_api/controller"
	middleware "github.com/fooksupachai/golang_restful_api/middleware"
)

func init() {

	router := mux.NewRouter()

	router.HandleFunc("/create_account", controller.CreateAccount).Methods("POST")
	router.HandleFunc("/auth", controller.Auth).Methods("POST")

	router.HandleFunc("/users", middleware.JWTMiddleware(controller.GetUsers)).Methods("GET")
	router.HandleFunc("/user", controller.CreateUser).Methods("POST")
	router.HandleFunc("/user/{firstname}/{lastname}", controller.UpdateUser).Methods("PUT")
	router.HandleFunc("/user/{firstname}", controller.DeleteUser).Methods("DELETE")
	router.HandleFunc("/user/{firstname}", controller.GetUser).Methods("GET")

	http.ListenAndServe(":8080", router)
}
