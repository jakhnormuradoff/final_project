package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/jakhnormuradoff/final_project/pkg/db"
)

func doneTaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "POST":
		id := r.URL.Query().Get("id")
		if id == "" {
			json.NewEncoder(w).Encode(map[string]string{"error": "ID is required"})
			return
		}
		task, err := db.TaskByID(id)
		if err != nil {
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to get task"})
			return
		}
		if task.Repeat == "" {
			err = db.DeleteTask(id)
			if err != nil {
				json.NewEncoder(w).Encode(map[string]string{"error": "Failed to delete task"})
				return
			}
			json.NewEncoder(w).Encode(map[string]string{})

			return
		}
		nextDate, err := NextDate(time.Now(), task.Date, task.Repeat)
		if err != nil {
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to calculate next date"})
			return
		}
		
		err = db.UpdateDate(nextDate, id)
		if err != nil {
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to update task date"})
			return
		}	
		json.NewEncoder(w).Encode(map[string]string{})
	default:
		json.NewEncoder(w).Encode(map[string]string{"error": "Method not allowed"})
		return
	}
}
