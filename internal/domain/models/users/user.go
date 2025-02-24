package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`

	Username string `gorm:"size:50;unique;not null" json:"username"`
	Password string `gorm:"size:255;not null" json:"password"`
	Email    string `gorm:"size:100;unique;not null" json:"email"`
	UserRole string `gorm:"type:varchar(20);default:'common'" json:"user_role"`

	EmailVerified     bool   `gorm:"default:false" json:"is_verified"`
	VerificationToken string `gorm:"size:100" json:"verification_token"`

	PasswordResetToken string     `gorm:"size:100" json:"password_reset_token"`
	TokenExpiry        *time.Time `json:"token_expiry"`
	ResetTokenExpiry   *time.Time `json:"reset_token_expiry"`
}

func (b *User) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.New()
	return
}