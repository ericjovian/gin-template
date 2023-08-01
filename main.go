package main

import (
	"log"

	"git.garena.com/sea-labs-id/batch-05/gin-template/db"
	"git.garena.com/sea-labs-id/batch-05/gin-template/server"
)

func main() {
	err := db.Connect()
	if err != nil {
		log.Println("Failed to connect DB", err)
	}
	server.Init()
}
