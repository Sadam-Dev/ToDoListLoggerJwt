package models

import "time"

type Task struct {
	ID           int       `gorm:"primary_key"`
	Title        string    `json:"title" gorm:"size:255"`
	Description  string    `json:"description" gorm:"size:255"`
	UserFullName string    `json:"userFullName" gorm:"size:255"`
	IsCompleted  bool      `json:"isCompleted" gorm:"default:false"`
	IsDeleted    bool      `json:"isDeleted" gorm:"default:false"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	UserID       int       `json:"userId" gorm:"default:0"`
}
