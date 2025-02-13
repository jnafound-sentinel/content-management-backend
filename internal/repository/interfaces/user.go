package interfaces

import (
	models "jna-manager/internal/domain/models/users"
	"time"
)

type UserRepository interface {
	Create(user *models.User) error
	GetByID(id string) (*models.User, error)
	GetByName(username string) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	Update(user *models.User) error
	Delete(id string) error
	List(page, pageSize int) ([]models.User, int64, error)

	GetByVerificationToken(token string) (*models.User, error)
	SetEmailVerified(userID string) error
	SetVerificationToken(userID string, token string, expiry time.Time) error

	GetByResetToken(token string, currentTime time.Time) (*models.User, error)
	UpdatePassword(userID string, hashedPassword string) error
	ClearResetToken(userID string) error
	SetResetToken(userID string, token string, expiry time.Time) error
}
