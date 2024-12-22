package main

import (
	"encoding/json"
	"net/http"
)

// TODO Доделать
func HandlerCard(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(items)

	case http.MethodPost:
		var newItem Item
		if err := json.NewDecoder(r.Body).Decode(&newItem); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}
		newItem.ID = len(items) + 1
		items = append(items, newItem)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newItem)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func HandlerCards(w http.ResponseWriter, r *http.Request) {

}

func HandlerAuth(w http.ResponseWriter, r *http.Request) {

}

func HandlerCalculator(w http.ResponseWriter, r *http.Request) {

}
