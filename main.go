package main

import (
	"log"

	"github.com/jakhnormuradoff/final_project/pkg/db"
	"github.com/jakhnormuradoff/final_project/pkg/server"
)

func main() {

	err := db.Init("scheduler.db")
	if err != nil {
		log.Fatal("Init db error :", err)
	}
	defer db.Close()

	err = server.StartServer("web", "7540")
	if err != nil {
		log.Fatal("error in StrartServer function :", err)
	}

}
