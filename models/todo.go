package models

import (
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	ID        string
	Task      string
	Completed bool
	CreatedBy string
}
