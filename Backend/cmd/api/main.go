package main

import (
	"enrollapp/internal/db"
	server "enrollapp/internal/server"
	"fmt"
	"log"
)

func main() {
	db.ConnectDB()

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Hello, World!")
}
