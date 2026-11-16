package task

import "context"

type Repository interface {
	List(ctx context.Context) ([]Task, error)
	Get(ctx context.Context, id int) (Task, error)
	Create(ctx context.Context, t Task) (Task, error)
	Update(ctx context.Context, t Task) (Task, error)
	Delete(ctx context.Context, id int) error
}
