package config

import (
	"backend/db/sqlite/tasksmock"
	models "backend/db/tasks"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type DBConfig struct {
	Path string
	DB   *gorm.DB
}

func NewDBConfig() *DBConfig {
	return &DBConfig{}
}

func (c *DBConfig) SetupDB() (*gorm.DB, error) {
	var err error
	var db *gorm.DB
	c.Path = "./db/sqlite/tasksmock/mock.db"
	db, err = gorm.Open("sqlite3", c.Path)
	if err != nil {
		return nil, fmt.Errorf("error connecting to sqlite database: %w", err)
	}
	c.DB = db

	err = c.SyncDb()
	if err != nil {
		return nil, fmt.Errorf("error running migrations: %w", err)
	}

	err = tasksmock.InsertDB(c.DB)
	if err != nil {
		return nil, fmt.Errorf("error inserting mock data: %w", err)
	}

	return c.DB, nil
}

func (c *DBConfig) SyncDb() error {
	err := c.DB.AutoMigrate(&models.Task{}).Error
	if err != nil {
		return fmt.Errorf("error running migrations %v", err)
	}
	return nil
}
