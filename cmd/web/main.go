package main

import (
	"fmt"
	"hangman/package/src/handler"
	"net/http"
)

const port = ":8080"

func main() {
	handler.Fs()

	fmt.Println("(http://localhost:8080)", "Server started on port ", port)

	// Gère la route "/"
	handler.Page()

	// Gère la route "/game"

	http.ListenAndServe(":8080", nil)
}
