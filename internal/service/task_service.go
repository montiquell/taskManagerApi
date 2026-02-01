package service

import (
	"context"
	"newTaskManagerApi/internal/domain"
	"newTaskManagerApi/internal/ports"
)

type TaskService struct {
	repo ports.TaskRepository
}

func NewTaskService(repo ports.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) Create(ctx context.Context, title string) (*domain.Task, error) {
	return s.repo.Create(ctx, title)
}

func (s *TaskService) GetAllTasks(ctx context.Context) ([]*domain.Task, error) {
	return s.repo.GetAllTasks(ctx)
}

func (s *TaskService) UpdateTaskStatus(
	ctx context.Context,
	id string,
	status bool,
) (*domain.Task, error) {
	return s.repo.UpdateTaskStatus(ctx, id, status)
}

func (s *TaskService) DeleteTask(ctx context.Context, id string) error {
	return s.repo.DeleteTask(ctx, id)
}
