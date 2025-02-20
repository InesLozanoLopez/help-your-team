package services

import (
	models "backend/db/tasks"
)

type TasksService struct {
	repo models.TasksRepositoryInterface
}

func NewTasksService(repo models.TasksRepositoryInterface) *TasksService {
	return &TasksService{repo: repo}
}

func (s *TasksService) FindAll() ([]models.TaskList, error) {
	ts, err := s.repo.All()
	r := make([]models.TaskList, len(ts))
	if err != nil {
		return r, err
	}

	for i, t := range ts {
		r[i] = models.TaskList{
			Id:          t.Id,
			Name:        t.Name,
			Description: t.Description,
			Status:      t.Status,
			AssignedTo:  t.AssignedTo,
		}
	}
	return r, nil
}

func (s *TasksService) FindById(id int) (models.Task, error) {
	return s.repo.TaskById(id)
}
