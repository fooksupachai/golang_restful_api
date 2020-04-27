package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	database "github.com/fooksupachai/golang_restful_api/database"
	u "github.com/fooksupachai/golang_restful_api/utils"
	"github.com/gorilla/mux"
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

	param := mux.Vars(r)

	database.UpdataUserData(r.Body, param["firstname"])

	resp := struct {
		Status int `json:"status"`
	}{
		Status: http.StatusAccepted,
	}

	json.NewEncoder(w).Encode(resp)
}

// DeleteUser to delete user account
func DeleteUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")

	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	param := mux.Vars(r)

	database.DeleteUserData(param["firstname"])

	resp := struct {
		Status int `json:"status"`
	}{
		Status: http.StatusAccepted,
	}

	json.NewEncoder(w).Encode(resp)
}

// GetUser by firstname
func GetUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	var accounts []Account

	param := mux.Vars(r)

	result := database.GetAccountData(param["firstname"])

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

	if accounts != nil {

		resp := struct {
			Account []Account `json:"account"`
			Status  int       `json:"status"`
		}{
			Account: accounts,
			Status:  http.StatusAccepted,
		}

		json.NewEncoder(w).Encode(resp)

	} else {

		resp := struct {
			Account string `json:"account"`
			Status  int    `json:"status"`
		}{
			Account: "Not found any account",
			Status:  http.StatusAccepted,
		}

		json.NewEncoder(w).Encode(resp)
	}
}

// UserConvert for provice modifer of user
func UserConvert(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	keyFirstname, ok := r.URL.Query()["firstname"]

	if !ok || len(keyFirstname[0]) < 1 {

		res := u.Message("error")
		res["data"] = "Missing firstname"
		u.Response(w, res)
		return

	}

	firstName := keyFirstname[0]

	fmt.Println(firstName)

	result := database.GetAccountData(firstName)

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

	if accounts != nil {

		resp := struct {
			Account []Account `json:"accounts"`
			Status  int       `json:"status"`
		}{
			Account: accounts,
			Status:  http.StatusAccepted,
		}

		json.NewEncoder(w).Encode(resp)
	} else {

		resp := struct {
			Account string `json:"account"`
			Status  int    `json:"status"`
		}{
			Account: "Not found any account",
			Status:  http.StatusAccepted,
		}

		json.NewEncoder(w).Encode(resp)
	}

}
