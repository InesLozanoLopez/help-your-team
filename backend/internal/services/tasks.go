package services

import (
	models "backend/db/tasks"
	"github.com/jinzhu/gorm"
)

type TasksService struct {
	repo *models.TasksRepository
}

func NewTasksService(db *gorm.DB) *TasksService {
	repo := models.NewTasksRepository(db)
	return &TasksService{repo: repo}
}

func (s *TasksService) FindAll() ([]models.TaskList, error) {
	var r []models.TaskList

	ts, err := s.repo.All()
	if err != nil {
		return r, err
	}
	for _, t := range ts {
		tr := models.TaskList{
			Id:          t.Id,
			Name:        t.Name,
			Description: t.Description,
			Status:      t.Status,
			AssignedTo:  t.AssignedTo,
		}
		r = append(r, tr)
	}
	return r, nil
}

func (s *TasksService) FindById(id int) (models.Task, error) {
	return s.repo.TaskById(id)
}
