package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"to-do-list-app/dto"
	"to-do-list-app/service"
)

type TaskHandler struct {
	taskService *service.TaskService
}

func NewTaskHandler(taskService *service.TaskService) *TaskHandler {
	return &TaskHandler{taskService: taskService}
}

// Create Task
func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var req dto.TaskCreateRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	taskResponse, err := h.taskService.CreateTask(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(taskResponse)
}

// Get All Tasks
func (h *TaskHandler) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.taskService.GetAllTasks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

// Delete Task
func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	// Ambil ID dari path ("/tasks/delete/{id}")
	segments := strings.Split(r.URL.Path, "/")
	if len(segments) < 3 {
		http.Error(w, "Invalid request URL", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(segments[3]) // Ambil ID dari path
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	err = h.taskService.SoftDeleteTask(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Task successfully deleted"})
}
