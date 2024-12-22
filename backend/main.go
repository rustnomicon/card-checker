package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/card/", HandlerCard)
	http.Handle("/cards/", authMiddleware(http.HandlerFunc(HandlerCards)))
	http.HandleFunc("/calculator/", HandlerCalculator)

	log.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
