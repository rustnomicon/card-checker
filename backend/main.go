package main

import (
	"log"
	"net/http"
)

func main() {
	db_connect := connect()
	defer db_connect.Close()
	// post - check valid card
	http.HandleFunc("/card/", HandlerCard(db_connect))
	// get all cards
	http.Handle("/cards/", authMiddleware(http.HandlerFunc(HandlerCards(db_connect)), db_connect))
	// get test calcultor
	http.HandleFunc("/calculator/", HandlerCalculator)

	log.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
