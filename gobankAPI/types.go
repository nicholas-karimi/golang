package main

import (
	"math/rand"
	"time"
)

type createAccountRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type Account struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Number    int64     `json:"number"`
	Balance   int64     `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}

func NewAccount(firstName string, lastName string) *Account {
	return &Account{
		// ID:        rand.Intn(1000),
		FirstName: firstName,
		LastName:  lastName,
		Number:    int64(rand.Intn(1000000000)),
		CreatedAt: time.Now(),
	}
}
