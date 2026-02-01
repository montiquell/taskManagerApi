package main

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

type TaskService struct {
	repo TaskRepository
}

func NewTaskService(repo TaskRepository) *TaskService {
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
