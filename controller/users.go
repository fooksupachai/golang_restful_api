package controller

import (
	"encoding/json"
	"net/http"

	database "github.com/fooksupachai/golang_restful_api/database"
)

// Account for describe infomation
type Account struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Age       int    `json:"age"`
	Address   string `json:"address"`
}

// GetUsers to Get all user in database
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	database.InitialDB()

	var account []Account

	accountList := Account{"Fook", "Fik", 23, "Detail"}

	for i := 0; i < 10; i++ {
		account = append(account, accountList)
	}

	resp := struct {
		Account []Account `json:"account"`
	}{
		Account: account,
	}

	json.NewEncoder(w).Encode(resp)

}
