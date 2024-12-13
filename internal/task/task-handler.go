package task

import (
	"encoding/json"
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

	err = s.CreateTask(r.Context(), taskDTO)
	if err != nil {
		http.Error(w, "Failed to create task", http.StatusInternalServerError)
		log.Printf("Error creating task: %s", err)
		return
	}

	// Отправляем успешный ответ с созданной задачей
	w.WriteHeader(http.StatusCreated)
	//json.NewEncoder(w).Encode(task)
}
