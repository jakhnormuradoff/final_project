package server
import (
	"fmt"
	"net/http"
)


func StartServer(webDir string, port string) error {
	fmt.Println("Запускаем сервер")
	http.Handle("/", http.FileServer(http.Dir(webDir)))
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		return err
	}
	fmt.Println("Завершаем работу")
	return nil
}