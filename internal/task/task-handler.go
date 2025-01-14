package task

import (
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"net/http"
)

func (s *Service) CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusBadRequest)
	}

	var taskDTO CreateTaskDTO
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&taskDTO)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		log.Printf("Error decoding task CreateTaskDTO: %s", err)
		return
	}
	var task *Task
	task, err = s.CreateTask(r.Context(), &taskDTO)
	if err != nil {
		http.Error(w, "Failed to create task", http.StatusInternalServerError)
		log.Printf("Error creating task: %s", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(task)
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}
}

func (s *Service) DeleteTaskByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method Not Allowed", http.StatusBadRequest)
	}
	var payload struct {
		Id string `json:"id"`
	}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&payload)
	if err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}

	fmt.Println(payload.Id, "payload.Id")

	err = s.DeleteTask(r.Context(), payload.Id)
	if err != nil {
		http.Error(w, "failed to delete task", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (s *Service) ChangeTaskByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method Not Allowed", http.StatusBadRequest)
	}
	var payload ChangeTaskDTO

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&payload)
	if err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}

	fmt.Println(payload.ID, "payload.Id", payload.Name, "payload.Name")

	changedTask, err := s.ChangeTask(r.Context(), payload)
	fmt.Println(changedTask)
	if err != nil {
		http.Error(w, "failed update task name", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(changedTask)
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}
}

func (s *Service) GetTaskByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusBadRequest)
	}
	id := r.URL.Query().Get("id")

	task, err := s.GetTask(r.Context(), id)
	slog.Info("get task by ID", "id", id, "task", task)
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(task)
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}
}
