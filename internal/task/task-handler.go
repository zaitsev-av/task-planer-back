package task

import (
	"encoding/json"
	"fmt"
	"log"
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

func (s *Service) DeleteTaskById(w http.ResponseWriter, r *http.Request) {
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

func (s *Service) ChangeTaskById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method Not Allowed", http.StatusBadRequest)
	}
	var payload struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&payload)
	if err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}

	fmt.Println(payload.Id, "payload.Id", payload.Name, "payload.Name")
	var changedTask *ChangeNameDTO
	changedTask, err = s.ChangeTask(r.Context(), payload.Id, payload.Name)
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
