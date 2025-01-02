package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"task-planer-back/config"
	ts "task-planer-back/internal/task"
	tr "task-planer-back/internal/task/db"
	"task-planer-back/pkg/client/postgresql"
	"task-planer-back/pkg/logger"
)

func main() {
	options := &slog.HandlerOptions{Level: slog.LevelInfo, AddSource: true}
	defaultHandler := slog.NewTextHandler(os.Stderr, options)
	colorHandler := logger.NewLoggerHandler(defaultHandler, options)

	customLogger := slog.New(colorHandler)
	slog.SetDefault(customLogger)

	cnf := config.GetConfig()

	client, err := postgresql.NewClient(context.Background(), cnf.Storage)
	if err != nil {
		slog.Error("Fatal err to connect db", "error", err)
		return
	}
	repo := tr.NewRepository(client)
	services := ts.NewServices(repo)

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/task", services.CreateTaskHandler)
	http.HandleFunc("/task/delete", services.DeleteTaskByID)
	http.HandleFunc("/task/change", services.ChangeTaskByID)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		slog.Info("Defaulting to port", port)
	}

	slog.Info("Listening on port", port)
	slog.Info("Open in the browser", "http://localhost:", port)
	server := &http.Server{
		Addr:              port,
		ReadHeaderTimeout: 5 * time.Second,
	}
	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
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
