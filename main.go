package main

import (
	"fmt"
	"log"
)

func main() {

	store, err := NewPostgresStore() // it return PostgresStore , and we method on it so that can Satisfied interface

	if err != nil {
		log.Fatal(err)

	}

	if err := store.Init(); err != nil {
		fmt.Println("after Init", err)

		log.Fatal(err)
	}
	server := NewAPIServer(":3000", store) // first argument string, second that can Satisfied out Storage interface
	server.Run()

}
