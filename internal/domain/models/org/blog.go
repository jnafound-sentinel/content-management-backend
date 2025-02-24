package models

import (
	"github.com/google/uuid"
)

type Blog struct {
	ID uuid.UUID `gorm:"type:uuid;primary_key;" json:"id"`

	Text string `gorm:"type:text" json:"text"`
}

func (e *Blog) TableName() string {
	return "blogs"
}
