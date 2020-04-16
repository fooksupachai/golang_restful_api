package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// InitialDB for connected to databases
func InitialDB() {
	db, err := sql.Open("mysql", "root:fyhwek+1t(aE@/book_store")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
}

// InsertData to database
func InsertData() {
	db, err := sql.Open("mysql", "root:fyhwek+1t(aE@/book_store")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	insert, err := db.Query("INSERT INTO Accounts VALUES ( 'Supachai', 'Keenthing', 23, '171/674 ...' )")

	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer insert.Close()
}
