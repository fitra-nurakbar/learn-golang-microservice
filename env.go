package main

import (
	"os"
)

func env() {
	os.Setenv("DATABASE", "development:mypassword@tcp(127.0.0.1:3306)/test")
}
