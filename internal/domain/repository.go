package domain

import "context"

type TaskRepository interface {
	Create(ctx context.Context, title string) (*Task, error)
	GetAllTasks(ctx context.Context) ([]*Task, error)
	UpdateTaskStatus(ctx context.Context, ID string, status bool) (*Task, error)
	DeleteTask(ctx context.Context, ID string) error
}
