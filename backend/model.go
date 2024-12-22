package main

type Card struct {
	ID          string `json:"id"`
	CardNumbers string `json:"numbers"`
	DateTime    string `json:"time"`
}

type Cards struct {
	Cards []Card `json"cards"`
}
