package models

import (
	"github.com/jinzhu/gorm"
)

type TasksRepository struct {
	Db *gorm.DB
}

func NewTasksRepository(db *gorm.DB) *TasksRepository {
	return &TasksRepository{Db: db}
}

func (r *TasksRepository) All() ([]Task, error) {
	var tasks []Task
	result := r.Db.Find(&tasks)
	return tasks, result.Error
}

func (r *TasksRepository) TaskById(id int) (Task, error) {
	var task Task
	result := r.Db.First(&task, id)
	return task, result.Error
}
