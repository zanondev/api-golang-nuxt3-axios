package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type Task struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var tasks []Task

func main() {
	http.HandleFunc("/tasks", getTasks)
	http.HandleFunc("/tasks/add", addTask)
	http.HandleFunc("/tasks/update", updateTask)

	http.ListenAndServe(":8080", nil)
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func addTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var newTask Task
	newTask.ID = uuid.New().String()

	err := json.NewDecoder(r.Body).Decode(&newTask)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tasks = append(tasks, newTask)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newTask)
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var updatedTask Task

	err := json.NewDecoder(r.Body).Decode(&updatedTask)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i, task := range tasks {
		if task.ID == updatedTask.ID {
			tasks[i] = updatedTask
			break
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(updatedTask)
	}
}
