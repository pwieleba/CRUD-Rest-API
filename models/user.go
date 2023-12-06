package models

type CreditCard struct {
	ID      string `json:"id"`
	Balance int    `json:"balance"`
	Owner   *Owner `json:"owner"`
}

type Owner struct {
	firstname string
	lastname  string
	age       int
	address   string
}


