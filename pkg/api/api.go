package api

import (
	"net/http"

	"github.com/jakhnormuradoff/final_project/pkg/config"
)

var appConfig *config.Config

func Init() {
	appConfig = config.LoadConfig()

	http.HandleFunc("/api/nextdate", NextDateHandler)
	http.HandleFunc("/api/task", auth(taskRouter))
	http.HandleFunc("/api/tasks", auth(tasksHandler))
	http.HandleFunc("/api/task/done", auth(doneTaskHandler))
	http.HandleFunc("/api/signin", signInHandler)
}
