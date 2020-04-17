package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

	"log"
	"os"
)

func getEnvVariable(key string) string {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv(key)
}

// InitialDB for connected to databases
func InitialDB() {

	password := getEnvVariable("PASSWORD_DATABASE")
	database := getEnvVariable("DATABASE_NAME")

	db, err := sql.Open("mysql", "root:"+password+database)

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
