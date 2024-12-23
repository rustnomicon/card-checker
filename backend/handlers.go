package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// TODO Доделать
func HandlerCard(db_connect *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			var card Card
			var cardNumbers string
			if err := json.NewDecoder(r.Body).Decode(&cardNumbers); err != nil {
				http.Error(w, "Invalid input", http.StatusBadRequest)
				return
			}

			timeNow := time.Now()

			card = Card{
				CardNumbers: cardNumbers,
				DateTime:    timeNow.String(),
			}

			id := insertCard(db_connect, card)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			response := fmt.Sprintf("Card id: %d", id)
			w.Write([]byte(response))

		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

func HandlerCards(db_connect *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func HandlerCalculator(w http.ResponseWriter, r *http.Request) {

}
