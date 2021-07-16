package models

import (
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	ID        string `json:"id"`
	Task      string `json:"task"`
	Completed bool   `json:"completed"`
	CreatedBy string `json:"created_by"`
}
