package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"

	_ "github.com/go-sql-driver/mysql"

	u "github.com/fooksupachai/golang_restful_api/utils"
)

// Account for describe infomation
type Account struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Age       int    `json:"age"`
	Address   string `json:"address"`
}

// Client for user handler
type Client struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Risk for risk handler
type Risk struct {
	Manufactory float64 `json:"manufactory"`
	Person      float64 `json:"person"`
	Government  float64 `json:"government"`
}

// InitialDB for connected to databases
func InitialDB() *sql.DB {

	password := u.GetEnvVariable(`PASSWORD_DATABASE`)
	database := u.GetEnvVariable(`DATABASE_NAME`)

	db, err := sql.Open(`mysql`, `root:`+password+database)

	if err != nil {
		panic(err.Error())
	}

	return db
}

// GetAllAccount user account from database accounts table
func GetAllAccount() *sql.Rows {

	db := InitialDB()

	result, err := db.Query(`SELECT * FROM book_store.Accounts`)

	if err != nil {
		panic(err.Error())
	}

	return result

}

// InsertData to database
func InsertData(body io.ReadCloser) {

	var account Account

	err := json.NewDecoder(body).Decode(&account)

	if err != nil {
		panic(err.Error())
	}

	db := InitialDB()

	insert, err := db.Prepare(`INSERT INTO Accounts VALUES (?, ?, ?, ?)`)

	_, err = insert.Exec(
		account.FirstName,
		account.LastName,
		account.Age,
		account.Address)

	if err != nil {
		panic(err.Error())
	}

	// be careful deferring Queries if you are using transactions
	defer insert.Close()
}

// UpdataUserData into database
func UpdataUserData(body io.ReadCloser, firstname string) {

	var account Account

	err := json.NewDecoder(body).Decode(&account)

	if err != nil {
		panic(err.Error())
	}

	db := InitialDB()

	update, err := db.Prepare(`UPDATE Accounts SET firstname = ?, lastname = ?, age = ?, address = ? WHERE firstname = ?`)

	_, err = update.Exec(
		account.FirstName,
		account.LastName,
		account.Age,
		account.Address,
		firstname,
	)

	if err != nil {
		panic(err.Error())
	}

	defer update.Close()
}

// DeleteUserData from database
func DeleteUserData(firstname string) {

	db := InitialDB()

	delete, err := db.Prepare(`DELETE FROM Accounts WHERE firstname = ?`)

	_, err = delete.Exec(
		firstname,
	)

	if err != nil {
		panic(err.Error())
	}

	defer delete.Close()
}

// GetAccountData from database
func GetAccountData(firstname string) *sql.Rows {

	db := InitialDB()

	find, err := db.Query(`Select * FROM Accounts WHERE firstname = ?`, firstname)

	if err != nil {
		panic(err.Error())
	}

	return find
}

// CreateAccontData into database
func CreateAccontData(Username string, Password string) {

	db := InitialDB()

	insert, err := db.Prepare(`INSERT INTO Client VALUES (?, ?)`)

	fmt.Println(Username, Password)
	_, err = insert.Exec(
		Username,
		Password,
	)

	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()
}

// CheckAccount to check account exist in database
func CheckAccount(body io.ReadCloser) *sql.Rows {

	var client Client

	json.NewDecoder(body).Decode(&client)

	db := InitialDB()

	find, err := db.Query(`SELECT * FROM Client WHERE Username = ? AND Password = ?`, client.Username, client.Password)

	fmt.Println(find)

	if err != nil {
		panic(err.Error())
	}

	return find
}

// InsertRiskData to database
func InsertRiskData(M float64, P float64, G float64) {

	db := InitialDB()

	insert, err := db.Prepare(`INSERT INTO Risk VALUES (?, ?, ?)`)

	fmt.Println("Database do")
	fmt.Println(M, P, G)

	_, err = insert.Exec(
		M,
		P,
		G,
	)

	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()

}

// GetAllRiskData from database
func GetAllRiskData() *sql.Rows {

	db := InitialDB()

	result, err := db.Query(`SELECT * FROM book_store.Risk`)

	if err != nil {
		panic(err.Error())
	}

	return result

}
