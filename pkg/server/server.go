package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jakhnormuradoff/final_project/pkg/api"
)

func StartServer(webDir string, port string) error {
	fmt.Println("Запускаем сервеsр")
	api.Init()
	http.Handle("/", http.FileServer(http.Dir(webDir)))
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Println("error in ListenAndServe:", err)
		return err
	}
	return nil
}
