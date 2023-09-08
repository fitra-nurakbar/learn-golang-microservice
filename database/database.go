package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var Query *sql.DB

func Connection() {
	var err error
	Query, err = sql.Open("mysql", os.Getenv("DATABASE"))

	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		os.Exit(1)
	}
}
