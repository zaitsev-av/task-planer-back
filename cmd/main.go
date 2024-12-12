package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"task-planer-back/pkg/client/postgresql"
)

func main() {
	getEnv()
	config := postgresql.Config{
		Username: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		Database: os.Getenv("POSTGRES_DB"),
	}
	data := os.Getenv("POSTGRES_HOST")
	fmt.Println(data, "POSTGRES_HOST")

	client, err := postgresql.NewClient(context.Background(), config)

	if err != nil {
		log.Fatal(err)
		return
	}
	if err := client.Ping(context.Background()); err != nil {
		log.Fatal("DB working", err)
		return
	}
	http.HandleFunc("/", indexHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Printf("Open http://localhost:%s in the browser", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
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

func getEnv() {
	err := godotenv.Load("../.env.docker")
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}
}
