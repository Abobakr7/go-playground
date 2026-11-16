package task

import (
	"context"
	"time"
)

type Service struct {
	repo Repository
}

func New(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) List(ctx context.Context) ([]Task, error) {
	return s.repo.List(ctx)
}

func (s *Service) Get(ctx context.Context, id int) (Task, error) {
	return s.repo.Get(ctx, id)
}

func (s *Service) Create(ctx context.Context, in Input) (Task, error) {
	if err := validate(&in); err != nil {
		return Task{}, err
	}

	return s.repo.Create(ctx, Task{
		Title:     in.Title,
		Done:      in.Done,
		CreatedAt: time.Now().UTC(),
	})
}

func (s *Service) Update(ctx context.Context, id int, in Input) (Task, error) {
	if err := validate(&in); err != nil {
		return Task{}, err
	}

	t, err := s.repo.Get(ctx, id)
	if err != nil {
		return Task{}, err
	}

	t.Title = in.Title
	t.Done = in.Done

	return s.repo.Update(ctx, t)
}

func (s *Service) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

func validate(in *Input) error {
	if in.Title == "" {
		return &ValidationError{Field: "title", Message: "is required"}
	}
	return nil
}
