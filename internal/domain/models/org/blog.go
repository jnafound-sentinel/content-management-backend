package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Blog struct {
	ID      uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Title   string    `gorm:"type:text" json:"title"`
	Excerpt string    `gorm:"type:text" json:"excerpt"`
	Date    time.Time `json:"date"`
	Tag     string    `gorm:"type:text" json:"tag"`
	Image   string    `gorm:"type:text" json:"image_url"`
	Content string    `gorm:"type:text" json:"content"`
}

func (b *Blog) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.New()
	return
}

func (e *Blog) TableName() string {
	return "blogs"
}
