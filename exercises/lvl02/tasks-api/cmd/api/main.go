package main

import (
	"log"
	"net/http"
	"tasks-api/internal/httpapi"
	"tasks-api/internal/task"
)

func main() {
	repo := task.NewMemoryRepository()
	svc := task.New(repo)
	api := httpapi.New(svc)

	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", api.Routes()))
}
