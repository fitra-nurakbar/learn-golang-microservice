package main

import (
	"fmt"
	"microservice/database"
	"net/http"
	"os"
)

func server() {
	env()
	database.Connection()
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/users", getAllUsers)
	http.HandleFunc("/user", getUserWithID)

	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
	os.Clearenv()
}
