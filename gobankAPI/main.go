package main

import (
	"log"
)

func main() {
	store, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("%+v\n", store)
	server := NewAPIServer(":8080", store)
	server.Run()
	// fmt.Println("Hello World")
}
