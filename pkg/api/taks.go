package api

import (
	"encoding/json"
	"net/http"

	"github.com/jakhnormuradoff/final_project/pkg/db"
)

type Tasks struct {
		Tasks []db.Task `json:"tasks"`
	}

func taskHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "POST":
		addTaskHandler(w, r)
	default:
		json.NewEncoder(w).Encode(map[string]string{"error": "Method not allowed"})
		return
	}
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		tasks, err := db.Tasks(50)
		if err != nil {
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to get tasks"})
			return
		}
		if tasks == nil {	
		tasks = []db.Task{}
	}
		json.NewEncoder(w).Encode(Tasks{Tasks: tasks})
	
	default:
		json.NewEncoder(w).Encode(map[string]string{"error": "Method not allowed"})
		return
	}
}
