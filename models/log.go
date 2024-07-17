package models

import (
	"time"

	"gorm.io/gorm"
)

type Log struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Timestamp   time.Time      `json:"timestamp"`
	Severity    string         `json:"severity"`
	ServiceName string         `json:"service_name"`
	Message     string         `json:"message"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
