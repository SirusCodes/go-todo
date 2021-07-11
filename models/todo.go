package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	ID        uuid.UUID
	Task      string
	Completed bool
	CreatedBy string
}
