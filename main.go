package main

import (
	"log"

	"github.com/ericjovian/gin-template/db"
	"github.com/ericjovian/gin-template/server"
)

func main() {
	err := db.Connect()
	if err != nil {
		log.Println("Failed to connect DB", err)
	}
	server.Init()
}
