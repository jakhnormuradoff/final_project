package api

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/jakhnormuradoff/final_project/pkg/db"
)

type TasksResponse struct {
	Tasks []db.Task `json:"tasks"`
}

const limit = 50

func getTaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query().Get("id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "ID is required"})
		return
	}
	task, err := db.TaskByID(id)

	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Task not found"})
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to get task"})
		return
	}

	json.NewEncoder(w).Encode(task)
}

func addTaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var task db.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to decode request body"})
		return
	}

	if task.Title == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Title is required"})
		return
	}

	now := time.Now()

	if task.Date == "" {
		task.Date = now.Format("20060102")
	} else {
		_, err := time.Parse("20060102", task.Date)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid date format"})
			return
		}
	}
	var next string
	if task.Repeat != "" {
		next, err = NextDate(now, task.Date, task.Repeat)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
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
		w.WriteHeader(http.StatusInternalServerError)
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
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to decode request body"})
		return
	}

	if task.ID == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "ID is required"})
		return
	}

	if task.Title == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Title is required"})
		return
	}

	now := time.Now()
	if task.Date == "" {
		task.Date = now.Format("20060102")
	} else {
		_, err := time.Parse("20060102", task.Date)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid date format"})
			return
		}
	}
	var next string
	if task.Repeat != "" {
		next, err = NextDate(now, task.Date, task.Repeat)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
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
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to update task"})
		return
	}

	json.NewEncoder(w).Encode(map[string]any{})
}

func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
	w.WriteHeader(http.StatusMethodNotAllowed)
	json.NewEncoder(w).Encode(map[string]string{"error": "Method not allowed"})
	return
}
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query().Get("id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "ID is required"})
		return
	}
	err := db.DeleteTask(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to delete task"})
		return
	}
	json.NewEncoder(w).Encode(map[string]string{})
}

func getTasksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tasks, err := db.Tasks(limit)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to get tasks"})
		return
	}
	if tasks == nil {
		tasks = []db.Task{}
	}
	json.NewEncoder(w).Encode(TasksResponse{Tasks: tasks})
}
