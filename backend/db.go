package main

import (
	"database/sql"
	"fmt"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "admin"
	password = "admin"
	dbname   = "testerdatabase"
)

func connect() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}

func insertCard(db *sql.DB, card Card) int {
	sqlStatement := `INSERT INTO cards (numbers, time) VALUES ($1, $2) RETURNING id`

	var id int
	err := db.QueryRow(sqlStatement, card.CardNumbers, card.DateTime).Scan(&id)
	if err != nil {
		panic(err)
	}

	return id
}

func getCards(db *sql.DB) ([]Card, error) {
	sqlStatement := `SELECT id, numbers, time from cards`

	rows, err := db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var cards []Card
	for rows.Next() {
		var card Card
		if err := rows.Scan(&card.ID, &card.CardNumbers, &card.DateTime); err != nil {
			return cards, err
		}
		cards = append(cards, card)
	}

	return cards, nil
}

func isAuthUser(db *sql.DB, sha string) bool {
	sqlStatement := `SELECT * FROM cards_auth where sha = $1`

	var id int
	err := db.QueryRow(sqlStatement, sha).Scan(&id)

	if err != nil {
		if err == sql.ErrNoRows {
			return false
		}
		panic(err)
	}

	return true
}
