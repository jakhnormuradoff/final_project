package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/jakhnormuradoff/final_project/pkg/db"
)

func addTaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var task db.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to decode request body"})
		return
	}

	if task.Title == "" {
		json.NewEncoder(w).Encode(map[string]string{"error": "Title is required"})
		return
	}

	now := time.Now()

	if task.Date == "" {
		task.Date = now.Format("20060102")
	} else {
		_, err := time.Parse("20060102", task.Date)
		if err != nil {
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid date format"})
			return
		}
	}
	var next string
	if task.Repeat != "" {
		next, err = NextDate(now, task.Date, task.Repeat)
		if err != nil {
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid repeat format"})
			return
		}
	}

	if task.Date < now.Format("20060102") {
		if task.Repeat == "" {
			task.Date = now.Format("20060102")
		} else {
			task.Date = next
		}
	}
	id, err := db.AddTask(&task)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to add task"})
		return
	}

	json.NewEncoder(w).Encode(map[string]int64{"id": id})

}


func updateTaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var task db.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to decode request body"})
		return
	}

	if task.ID == "" {
		json.NewEncoder(w).Encode(map[string]string{"error": "ID is required"})
		return
	}

	if task.Title == "" {
		json.NewEncoder(w).Encode(map[string]string{"error": "Title is required"})
		return
	}

	now := time.Now()
	if task.Date == "" {
		task.Date = now.Format("20060102")
	} else {
		_, err := time.Parse("20060102", task.Date)
		if err != nil {
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid date format"})
			return
		}
	}
	var next string
	if task.Repeat != "" {
		next, err = NextDate(now, task.Date, task.Repeat)
		if err != nil {
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid repeat format"})
			return
		}
	}

	if task.Date < now.Format("20060102") {
		if task.Repeat == "" {
			task.Date = now.Format("20060102")
		} else {
			task.Date = next
		}
	}
	 err = db.UpdateTask(&task)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to update task"})
		return
	}

	json.NewEncoder(w).Encode(map[string]any{})
}	