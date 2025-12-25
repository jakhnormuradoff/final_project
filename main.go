package main

import (
	"fmt"

	"github.com/jakhnormuradoff/final_project/pkg/server"
)

func main() {

	err := server.StartServer("web", "7540")
	if err != nil {
		panic(err)

	}
	fmt.Println("Завершаем работу")
}
