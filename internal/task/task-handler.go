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

	var taskDTO DTO
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&taskDTO)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		log.Printf("Error decoding task DTO: %s", err)
		return
	}
	var task *Task
	task, err = s.CreateTask(r.Context(), &taskDTO)
	if err != nil {
		http.Error(w, "Failed to create task", http.StatusInternalServerError)
		log.Printf("Error creating task: %s", err)
		return
	}

	//responseData, err := json.Marshal(task)
	//if err != nil {
	//	http.Error(w, fmt.Sprintf("Failed to encode JSON: %v", err), http.StatusInternalServerError)
	//	return
	//}

	w.Header().Set("Content-Type", "application/json")
	// Отправляем успешный ответ с созданной задачей
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
