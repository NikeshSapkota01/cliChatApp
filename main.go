package main

import (
	"log"

	"github.com/NikeshSapkota01/cliChatApp/cmd"
	db "github.com/NikeshSapkota01/cliChatApp/db"
)

func main() {

	_, err := db.NewDatabase()

	if err != nil {
		log.Fatalf("Couldnot initialize the database connection: %s", err)
	}

	log.Printf("Connected to DB ...")

	cmd.Execute()
}
