package services

import (
	models "backend/db/tasks"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type MockTasksRepository struct{}

func (m *MockTasksRepository) Find(dest interface{}, _ ...interface{}) *gorm.DB {
	var tnil *time.Time
	mock := []models.Task{
		{Id: 1,
			Name:        "Create landing page",
			Description: "Design and create the main landing page of the website.",
			Status:      "To Do",
			AssignedTo:  "John Bush",
			CreatedAt:   time.Date(2025, 02, 01, 10, 00, 00, 00, time.UTC),
			CompletedAt: tnil,
			Category:    "Design",
		},
		{
			Id:          2,
			Name:        "Review bug in contact form",
			Description: "The contact form does not send messages correctly.",
			Status:      "In Progress",
			AssignedTo:  "Anna John",
			CreatedAt:   time.Date(2025, 02, 05, 12, 30, 0, 00, time.UTC),
			CompletedAt: tnil,
			Category:    "Bug",
		},
		{
			Id:          3,
			Name:        "Write content for blog",
			Description: "Write 5 articles about digital marketing.",
			Status:      "To Do",
			AssignedTo:  "Charles King",
			CreatedAt:   time.Date(2025, 02, 02, 9, 00, 00, 00, time.UTC),
			CompletedAt: tnil,
			Category:    "Content",
		},
		{
			Id:          4,
			Name:        "Implement user API",
			Description: "Develop the API for user management.",
			Status:      "In Progress",
			AssignedTo:  "David Stinton",
			CreatedAt:   time.Date(2025, 02, 10, 14, 00, 00, 00, time.UTC),
			CompletedAt: tnil,
			Category:    "Development",
		}}
	*dest.(*[]models.Task) = mock
	return &gorm.DB{}
}

func (m *MockTasksRepository) First(dest interface{}, args ...interface{}) *gorm.DB {
	if len(args) > 0 {
		id, ok := args[0].(int)
		if !ok {
			return &gorm.DB{}
		}
		var tnil *time.Time
		mock := []models.Task{
			{Id: 1,
				Name:        "Create landing page",
				Description: "Design and create the main landing page of the website.",
				Status:      "To Do",
				AssignedTo:  "John Bush",
				CreatedAt:   time.Date(2025, 02, 01, 10, 00, 00, 00, time.UTC),
				CompletedAt: tnil,
				Category:    "Design",
			},
			{
				Id:          2,
				Name:        "Review bug in contact form",
				Description: "The contact form does not send messages correctly.",
				Status:      "In Progress",
				AssignedTo:  "Anna John",
				CreatedAt:   time.Date(2025, 02, 05, 12, 30, 0, 00, time.UTC),
				CompletedAt: tnil,
				Category:    "Bug",
			},
			{
				Id:          3,
				Name:        "Write content for blog",
				Description: "Write 5 articles about digital marketing.",
				Status:      "To Do",
				AssignedTo:  "Charles King",
				CreatedAt:   time.Date(2025, 02, 02, 9, 00, 00, 00, time.UTC),
				CompletedAt: tnil,
				Category:    "Content",
			},
			{
				Id:          4,
				Name:        "Implement user API",
				Description: "Develop the API for user management.",
				Status:      "In Progress",
				AssignedTo:  "David Stinton",
				CreatedAt:   time.Date(2025, 02, 10, 14, 00, 00, 00, time.UTC),
				CompletedAt: tnil,
				Category:    "Development",
			}}

		for _, t := range mock {
			if t.Id == uint(id) {
				*(dest.(*models.Task)) = t
				return &gorm.DB{}
			}
		}
	}
	return &gorm.DB{Error: gorm.ErrRecordNotFound}
}

func TestTasksService_FindAll(t *testing.T) {
	mockRepo := &MockTasksRepository{}
	service := models.NewTasksRepository(mockRepo)

	ts, err := service.All()
	assert.NoError(t, err)
	assert.Equal(t, 4, len(ts))
}

func TestTasksService_FindById(t *testing.T) {
	mockRepo := &MockTasksRepository{}
	service := models.NewTasksRepository(mockRepo)

	ta, err := service.TaskById(1)
	assert.NoError(t, err)
	assert.Equal(t, uint(1), ta.Id)

	_, err = service.TaskById(22)
	assert.Error(t, err)
	assert.ErrorIs(t, err, gorm.ErrRecordNotFound)
}
