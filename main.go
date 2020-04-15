package main

import (
	"log"
	"net/http"

	_ "github.com/fooksupachai/golang_restful_api/router"
)

func main() {
	log.Print("Rest server is listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
