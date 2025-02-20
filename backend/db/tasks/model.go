package models

import "time"

type Task struct {
	Id          uint      `gorm:"primary_key" json:"id"`
	Name        string    `gorm:"type:varchar(100);not null" json:"name"`
	Description string    `gorm:"type:text;not null" json:"description"`
	Status      string    `gorm:"type:varchar(20);not null" json:"status"`
	AssignedTo  string    `gorm:"type:varchar(100);" json:"assigned_to,omitempty"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP;not null" json:"created_at"`
	CompletedAt time.Time `gorm:"type:timestamp;" json:"completed_at,omitempty"`
	Category    string    `gorm:"type:varchar(20);not null" json:"category"`
}

type TaskList struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
	AssignedTo  string `json:"assigned_to,omitempty"`
}
