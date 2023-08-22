package main

import (
	"fmt"
	"net/http"
	"os"
)

func server() {
	env()
	database()
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/login", loginHandler)

	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
	os.Clearenv()
}
