package api

import (
	"encoding/json"
	"net/http"
)

func taskRouter(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "POST":
		addTaskHandler(w, r)
	case "GET":
		getTaskHandler(w, r)
	case "PUT":
		updateTaskHandler(w, r)
	case "DELETE":
		deleteTaskHandler(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"error": "Method not allowed"})
		return
	}
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		getTasksHandler(w, r)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"error": "Method not allowed"})
		return
	}
}
