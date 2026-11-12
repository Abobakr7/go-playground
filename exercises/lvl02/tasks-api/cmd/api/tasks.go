package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sort"
	"strconv"
	"sync"
	"time"
)

type Task struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"created_at"`
}

type taskInput struct {
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

type server struct {
	mu     sync.RWMutex
	tasks  map[int]Task
	nextId int
}

func newServer() *server {
	return &server{tasks: make(map[int]Task), nextId: 1}
}

func (s *server) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", s.health)
	mux.HandleFunc("GET /tasks", s.listTasks)
	mux.HandleFunc("POST /tasks", s.createTask)
	mux.HandleFunc("GET /tasks/{id}", s.getTask)
	mux.HandleFunc("PUT /tasks/{id}", s.updateTask)
	mux.HandleFunc("DELETE /tasks/{id}", s.deleteTask)
	return mux
}

// ---- helpers ----

func writeJson(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Printf("encode response: %v", err)
	}
}

func writeError(w http.ResponseWriter, status int, msg string) {
	writeJson(w, status, map[string]string{"error": msg})
}

func parseID(r *http.Request) (int, bool) {
	id, err := strconv.Atoi(r.PathValue("id"))
	return id, err == nil && id > 0
}

// ---- handlers ----

func (s *server) health(w http.ResponseWriter, r *http.Request) {
	writeJson(w, http.StatusOK, map[string]string{"status": "ok"})
}

func (s *server) listTasks(w http.ResponseWriter, r *http.Request) {
	s.mu.RLock()
	out := make([]Task, 0, len(s.tasks))
	for _, t := range s.tasks {
		out = append(out, t)
	}
	s.mu.RUnlock()

	sort.Slice(out, func(i, j int) bool { return out[i].ID > out[j].ID })
	writeJson(w, http.StatusOK, out)
}

func (s *server) createTask(w http.ResponseWriter, r *http.Request) {
	var in taskInput
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}
	if in.Title == "" {
		writeError(w, http.StatusBadRequest, "title is required")
		return
	}

	s.mu.Lock()
	t := Task{ID: s.nextId, Title: in.Title, Done: in.Done, CreatedAt: time.Now().UTC()}
	s.tasks[s.nextId] = t
	s.nextId++
	s.mu.Unlock()

	w.Header().Set("Location", "/tasks/"+strconv.Itoa(t.ID))
	writeJson(w, http.StatusCreated, t)
}

func (s *server) getTask(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(r)
	if !ok {
		writeError(w, http.StatusBadRequest, "invalid task id")
		return
	}

	s.mu.RLock()
	t, found := s.tasks[id]
	s.mu.RUnlock()

	if !found {
		writeError(w, http.StatusNotFound, "task not found")
		return
	}

	writeJson(w, http.StatusOK, t)
}

func (s *server) updateTask(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(r)
	if !ok {
		writeError(w, http.StatusBadRequest, "invalid task id")
		return
	}

	var in taskInput
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}
	if in.Title == "" {
		writeError(w, http.StatusBadRequest, "title is required")
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	t, found := s.tasks[id]
	if !found {
		writeError(w, http.StatusNotFound, "task not found")
		return
	}

	t.Title = in.Title
	t.Done = in.Done
	s.tasks[id] = t

	writeJson(w, http.StatusOK, t)
}

func (s *server) deleteTask(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(r)
	if !ok {
		writeError(w, http.StatusBadRequest, "invalid task id")
		return
	}

	s.mu.Lock()
	_, found := s.tasks[id]
	delete(s.tasks, id)
	s.mu.Unlock()

	if !found {
		writeError(w, http.StatusNotFound, "task not found")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
