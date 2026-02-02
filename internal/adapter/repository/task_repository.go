package repository

import (
	"context"
	"newTaskManagerApi/internal/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TaskModel struct {
	ID        string `gorm:"primaryKey;type:uuid"`
	Title     string `gorm:"not null"`
	Completed bool   `gorm:"not null"`
}

func (TaskModel) TableName() string {
	return "tasks"
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *taskRepository {
	return &taskRepository{db: db}
}

func (repo *taskRepository) Create(ctx context.Context, title string) (*domain.Task, error) {
	model := &TaskModel{
		ID:        uuid.New().String(),
		Title:     title,
		Completed: false,
	}

	err := repo.db.WithContext(ctx).Create(model).Error
	if err == gorm.ErrDuplicatedKey {
		return nil, domain.ErrAlreadyExists
	} else if err != nil {
		return nil, err
	}

	return toDomain(model), nil
}

func (repo *taskRepository) GetAllTasks(ctx context.Context) ([]*domain.Task, error) {
	var models []TaskModel

	err := repo.db.WithContext(ctx).Find(&models).Error
	if err != nil {
		return nil, err
	}

	tasks := make([]*domain.Task, len(models))
	for idx, model := range models {
		tasks[idx] = toDomain(&model)
	}
	return tasks, nil
}

func (repo *taskRepository) UpdateTaskStatus(
	ctx context.Context, ID string, status bool,
) (*domain.Task, error) {
	var model TaskModel

	err := repo.db.WithContext(ctx).Where("id = ?", ID).First(&model).Error
	if err == gorm.ErrRecordNotFound {
		return nil, domain.ErrNotFound
	} else if err != nil {
		return nil, err
	}

	err = repo.db.WithContext(ctx).
		Model(&TaskModel{}).
		Where("id = ?", ID).
		Update("completed", status).
		Error

	if err != nil {
		return nil, err
	}

	model.Completed = status

	task := toDomain(&model)
	return task, nil
}

func (repo *taskRepository) DeleteTask(ctx context.Context, ID string) error {
	result := repo.db.WithContext(ctx).Where("id = ?", ID).Delete(&TaskModel{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return domain.ErrNotFound
	}

	return nil
}

func toDomain(model *TaskModel) *domain.Task {
	return &domain.Task{
		ID:        model.ID,
		Title:     model.Title,
		Completed: model.Completed,
	}
}
