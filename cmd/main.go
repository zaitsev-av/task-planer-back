package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"task-planer-back/config"
	taskServices "task-planer-back/internal/task"
	taskRepo "task-planer-back/internal/task/db"
	"task-planer-back/pkg/client/postgresql"
)

func main() {
	cnf := config.GetConfig()

	client, err := postgresql.NewClient(context.Background(), cnf.Storage)
	if err != nil {
		log.Fatal("Fatal err to connect db", err)
		return
	}
	repo := taskRepo.NewRepository(client)
	services := taskServices.NewServices(repo)
	//services.TaskServices(context.Context())

	if err != nil {
		log.Fatal(err)
		return
	}
	if err := client.Ping(context.Background()); err != nil {
		log.Fatal("DB not working", err)
		return
	}
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/task", services.CreateTaskHandler)
	http.HandleFunc("/task/delete", services.DeleteTaskById)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Printf("Open http://localhost:%s in the browser", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))

	//repo := db.NewRepository()
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	_, err := fmt.Fprint(w, "Hello, World!")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
