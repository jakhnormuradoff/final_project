package main

import (
	"fmt"
	"net/http"
	
)

func main() {
	fmt.Println("Запускаем сервер")
	err := StartServer("web", "7540")
	//http.Handle("/", http.FileServer(http.Dir(webDir)))
	err = http.ListenAndServe(":7540", nil)
	if err != nil {
		panic(err)

	}
	fmt.Println("Завершаем работу")
}
