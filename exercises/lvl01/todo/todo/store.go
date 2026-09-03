package todo

import "fmt"

type Store struct {
	todos []Todo
	nextId int
}

var store *Store = &Store{
	todos: make([]Todo, 0),
	nextId: 1,
}

func NewStore() *Store {
	return store
}

func (s *Store) Add(title string) Todo {
	todo := Todo{
		ID:     s.nextId,
		Title:  title,
		Status: Open,
	}
	s.todos = append(s.todos, todo)
	s.nextId++
	return todo
}

func (s *Store) List() []Todo {
	return s.todos
}

func (s *Store) Complete(id int) error {
	for i := range s.todos {
		if s.todos[i].ID == id {
			s.todos[i].Status = Done
			return nil
		}
	}
	return fmt.Errorf("Todo with ID %d not found", id)
}

func (s *Store) Delete(id int) error {
	for i := range s.todos {
		if s.todos[i].ID == id {
			s.todos = append(s.todos[:i], s.todos[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Todo with ID %d not found", id)
}
