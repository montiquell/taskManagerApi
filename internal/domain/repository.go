package domain

import "context"

type TaskRepository interface {
	Create(ctx context.Context, title string)
	GetAllTasks(ctx context.Context)
	UpdateTaskStatus(ctx context.Context, ID string, status bool)
	DeleteTask(ctx context.Context, ID string)
}
