package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// InitialDB for connected to databases
func InitialDB()  {
	db, err := sql.Open("mysql", "root:fyhwek+1t(aE@/book_store")
	defer db.Close()
	if err != nil {
		fmt.Println("connected failed")
	}else {
		fmt.Println("connected successed")
	}
}