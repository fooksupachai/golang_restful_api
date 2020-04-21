package router

import (
	"net/http"

	"github.com/gorilla/mux"

	controller "github.com/fooksupachai/golang_restful_api/controller"
)

func init() {

	router := mux.NewRouter()

	router.HandleFunc("/auth", controller.CreateAccount).Methods("POST")

	router.HandleFunc("/users", controller.GetUsers).Methods("GET")
	router.HandleFunc("/user", controller.CreateUser).Methods("POST")
	router.HandleFunc("/user/{firstname}/{lastname}", controller.UpdateUser).Methods("PUT")
	router.HandleFunc("/user/{firstname}", controller.DeleteUser).Methods("DELETE")
	router.HandleFunc("/user/{firstname}", controller.GetUser).Methods("GET")

	router.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	}).Methods("GET")

	http.ListenAndServe(":8080", router)
}
