package schemas

import (
	"time"

	"github.com/google/uuid"
)

type CreateBlogRequest struct {
	Title   string    `gorm:"type:text" json:"title" binding:"required,min=1,max=255"`
	Excerpt string    `gorm:"type:text" json:"excerpt" binding:"required,min=1,max=500"`
	Date    time.Time `json:"date" binding:"required"`
	Tag     string    `gorm:"type:text" json:"tag" binding:"required,min=1,max=100"`
	Image   string    `gorm:"type:text" json:"image_url" binding:"required,url"`
	Content string    `gorm:"type:text" json:"content" binding:"required,min=1"`
}

type UpdateBlogRequest struct {
	Title   string    `gorm:"type:text" json:"title" binding:"omitempty,min=1,max=255"`
	Excerpt string    `gorm:"type:text" json:"excerpt" binding:"omitempty,min=1,max=500"`
	Date    time.Time `json:"date" binding:"omitempty"`
	Tag     string    `gorm:"type:text" json:"tag" binding:"omitempty,min=1,max=100"`
	Image   string    `gorm:"type:text" json:"image_url" binding:"omitempty,url"`
	Content string    `gorm:"type:text" json:"content" binding:"omitempty,min=1"`
}

type BlogResponse struct {
	ID      uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Title   string    `gorm:"type:text" json:"title"`
	Excerpt string    `gorm:"type:text" json:"excerpt"`
	Date    time.Time `json:"date"`
	Tag     string    `gorm:"type:text" json:"tag"`
	Image   string    `gorm:"type:text" json:"image_url"`
	Content string    `gorm:"type:text" json:"content"`
}