package task

import (
	"errors"
	"time"
)

type Task struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"created_at"`
}

type Input struct {
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var ErrNotFound = errors.New("task not found")

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (e *ValidationError) Error() string {
	return e.Field + " " + e.Message
}
