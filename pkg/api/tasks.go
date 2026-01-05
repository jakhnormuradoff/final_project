package api

import (
	"database/sql"
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
	case "GET":
		id := r.URL.Query().Get("id")

		if id == "" {
			json.NewEncoder(w).Encode(map[string]string{"error": "ID is required"})
			return
		}
		task, err := db.TaskByID(id)

		if err == sql.ErrNoRows {
			json.NewEncoder(w).Encode(map[string]string{"error": "Task not found"})
			return
		}

		if err != nil {
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to get task"})
			return
		}

		json.NewEncoder(w).Encode(task)

	case "PUT":
		updateTaskHandler(w, r)
	case "DELETE":
		id := r.URL.Query().Get("id")
		if id == "" {
			json.NewEncoder(w).Encode(map[string]string{"error": "ID is required"})
			return
		}
		err := db.DeleteTask(id)
		if err != nil {
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to delete task"})
			return
		}
		json.NewEncoder(w).Encode(map[string]string{})
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
