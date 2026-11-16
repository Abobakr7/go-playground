package task

import (
	"context"
	"sort"
	"sync"
)

type MemoryRepository struct {
	mu     sync.RWMutex
	tasks  map[int]Task
	nextID int
}

var _ Repository = (*MemoryRepository)(nil)

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{tasks: make(map[int]Task), nextID: 1}
}

func (mr *MemoryRepository) List(ctx context.Context) ([]Task, error) {
	mr.mu.RLock()
	defer mr.mu.RUnlock()

	out := make([]Task, 0, len(mr.tasks))
	for _, t := range mr.tasks {
		out = append(out, t)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].ID < out[j].ID })

	return out, nil
}

func (mr *MemoryRepository) Get(ctx context.Context, id int) (Task, error) {
	mr.mu.RLock()
	defer mr.mu.RUnlock()

	t, found := mr.tasks[id]
	if !found {
		return Task{}, ErrNotFound
	}

	return t, nil
}

func (mr *MemoryRepository) Create(ctx context.Context, t Task) (Task, error) {
	mr.mu.Lock()
	defer mr.mu.Unlock()

	t.ID = mr.nextID
	mr.nextID++
	mr.tasks[t.ID] = t

	return t, nil
}

func (mr *MemoryRepository) Update(ctx context.Context, t Task) (Task, error) {
	mr.mu.Lock()
	defer mr.mu.Unlock()

	if _, found := mr.tasks[t.ID]; !found {
		return Task{}, ErrNotFound
	}

	mr.tasks[t.ID] = t
	return t, nil
}

func (mr *MemoryRepository) Delete(ctx context.Context, id int) error {
	mr.mu.Lock()
	defer mr.mu.Unlock()

	if _, found := mr.tasks[id]; !found {
		return ErrNotFound
	}

	delete(mr.tasks, id)
	return nil
}
