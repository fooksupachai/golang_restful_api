package router

import (
	"net/http"

	"github.com/gorilla/mux"

	controller "github.com/fooksupachai/golang_restful_api/controller"
)

func init() {

	router := mux.NewRouter()

	router.HandleFunc("/users", controller.GetUsers).Methods("GET")
	router.HandleFunc("/user", controller.CreateUser).Methods("POST")
	router.HandleFunc("/user/{firstname}/{lastname}", controller.UpdateUser).Methods("PUT")
	router.HandleFunc("/user/{firstname}", controller.DeleteUser).Methods("DELETE")
	router.HandleFunc("/user/{firstname}", controller.GetUser).Methods("GET")

	http.ListenAndServe(":8080", router)
}
