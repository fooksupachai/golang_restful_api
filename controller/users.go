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

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	result := database.GetAllAccount()

	var accounts []Account

	for result.Next() {

		var account Account
		err := result.Scan(
			&account.FirstName,
			&account.LastName,
			&account.Age,
			&account.Address,
		)

		if err != nil {
			panic(err.Error())
		}

		accounts = append(accounts, account)
	}

	resp := struct {
		Account []Account `json:"account"`
	}{
		Account: accounts,
	}

	json.NewEncoder(w).Encode(resp)

}

// CreateUser to generate user account
func CreateUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	database.InsertData(r.Body)

	resp := struct {
		Status int `json:"status"`
	}{
		Status: http.StatusAccepted,
	}

	json.NewEncoder(w).Encode(resp)
}
