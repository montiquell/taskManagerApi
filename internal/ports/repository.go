package ports

import (
	"context"
	"newTaskManagerApi/internal/domain"
)

type TaskRepository interface {
	Create(ctx context.Context, title string) (*domain.Task, error)
	GetAllTasks(ctx context.Context) ([]*domain.Task, error)
	UpdateTaskStatus(ctx context.Context, ID string, status bool) (*domain.Task, error)
	DeleteTask(ctx context.Context, ID string) error
}
