package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID string `gorm:"type:uuid;primary_key;"`

	Username string `gorm:"size:50;unique;not null"`
	Password string `gorm:"size:255;not null"`
	Email    string `gorm:"size:100;unique;not null"`
	UserRole string `gorm:"type:varchar(20);default:'student'"`

	EmailVerified     bool   `gorm:"default:false"`
	VerificationToken string `gorm:"size:100"`

	PasswordResetToken string `gorm:"size:100"`
	TokenExpiry        *time.Time
	ResetTokenExpiry   *time.Time
}

func (b *User) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.New().String()
	return
}
