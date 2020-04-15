package router

import (
	"net/http"

	controller "github.com/fooksupachai/golang_restful_api/controller"
)

func init() {
	http.HandleFunc("/users", controller.GetUsers)
}
