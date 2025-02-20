package services

import (
	"backend/db/sqlite/tasksmock"
	models "backend/db/tasks"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"testing"
)

func setupTestDB() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&models.Task{})
	return db, nil
}

func TestTasksService_FindAll(t *testing.T) {
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("Failed to set up test DB: %v", err)
	}
	err = tasksmock.InsertDB(db)
	if err != nil {
		t.Fatalf("error inserting mock data: %v", err)
	}

	service := NewTasksService(db)

	ts, err := service.FindAll()
	assert.NoError(t, err)
	assert.Equal(t, 20, len(ts))
}

func TestTasksService_FindById(t *testing.T) {
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("Failed to set up test DB: %v", err)
	}
	err = tasksmock.InsertDB(db)
	if err != nil {
		t.Fatalf("error inserting mock data: %v", err)
	}

	service := NewTasksService(db)

	ta, err := service.FindById(1)
	assert.NoError(t, err)
	assert.Equal(t, uint(1), ta.Id)

	_, err = service.FindById(22)
	assert.Error(t, err)
	assert.ErrorIs(t, err, gorm.ErrRecordNotFound)
}
