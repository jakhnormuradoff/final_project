package server

import (
	"fmt"
	"log"
	"net/http"
)


func StartServer(webDir string, port string) error {
	fmt.Println("Запускаем сервер")
	http.Handle("/", http.FileServer(http.Dir(webDir)))
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Println("error in ListenAndServe:", err)
		return err
	}
	return nil
}