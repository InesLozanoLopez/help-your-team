package models

import (
	"github.com/jinzhu/gorm"
)

type DbInterface interface {
	Find(dest interface{}, args ...interface{}) *gorm.DB
	First(dest interface{}, args ...interface{}) *gorm.DB
}
type TasksRepository struct {
	Db DbInterface
}

func NewTasksRepository(db DbInterface) *TasksRepository {
	return &TasksRepository{Db: db}
}

type TasksRepositoryInterface interface {
	All() ([]Task, error)
	TaskById(id int) (Task, error)
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
