package main

import (
	"fmt"
	"log"
	"net/http"

	"asciiart/internal/delivery"
)

func main() {
	server := delivery.New()
	fmt.Printf("Starting server at port :8080\nhttp://localhost:8080\n")
	log.Fatal(http.ListenAndServe(":8080", server.Router()))
}
