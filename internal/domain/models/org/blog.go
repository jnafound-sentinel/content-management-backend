package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Blog struct {
	ID uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`

	Text string `gorm:"type:text" json:"text"`
}

func (b *Blog) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.New()
	return
}

func (e *Blog) TableName() string {
	return "blogs"
}
