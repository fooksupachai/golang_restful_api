package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

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

// UpdateUser to update user describe
func UpdateUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")

	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	param := r.URL.Path[6:]

	idx := strings.Index(param, "/")
	if idx == -1 {

		// url /user/:id we're all set
		fmt.Println(param)
	}

	static := param[idx+1:]

	// found slash but no more to the URL
	if len(static) == 0 {
		http.Redirect(w, r, r.URL.Path[:len(r.URL.Path)-1], http.StatusMovedPermanently)
		return
	}

	database.UpdataUserData(r.Body, param)

	resp := struct {
		Status int `json:"status"`
	}{
		Status: http.StatusAccepted,
	}

	json.NewEncoder(w).Encode(resp)
}
