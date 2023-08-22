package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func database() {
	var err error
	db, err = sql.Open("mysql", os.Getenv("DATABASE"))

	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		os.Exit(1)
	}
}
