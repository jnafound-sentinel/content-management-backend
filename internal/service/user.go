package service

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	models "jna-manager/internal/domain/models/users"
	"jna-manager/internal/domain/schemas"
	"jna-manager/internal/repository/interfaces"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo interfaces.UserRepository
}

func NewUserService(repo interfaces.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	token := make([]byte, 32)
	if _, err := rand.Read(token); err != nil {
		return err
	}
	verificationToken := base64.URLEncoding.EncodeToString(token)

	expiry := time.Now().Add(24 * time.Hour)
	user.VerificationToken = verificationToken
	user.TokenExpiry = &expiry
	user.EmailVerified = false

	return s.repo.Create(user)
}

func (s *UserService) VerifyEmail(token string) error {
	user, err := s.repo.GetByVerificationToken(token)
	if err != nil {
		return err
	}

	if user.TokenExpiry.Before(time.Now()) {
		return errors.New("verification token has expired")
	}

	return s.repo.SetEmailVerified(user.ID)
}

func (s *UserService) ResendVerification(email string) (string, error) {
	user, err := s.repo.GetByEmail(email)
	if err != nil {
		return "", err
	}

	if user.EmailVerified {
		return "", errors.New("email is already verified")
	}

	token := make([]byte, 32)
	if _, err := rand.Read(token); err != nil {
		return "", err
	}
	verificationToken := base64.URLEncoding.EncodeToString(token)

	expiry := time.Now().Add(24 * time.Hour)

	if err := s.repo.SetVerificationToken(user.ID, verificationToken, expiry); err != nil {
		return "", err
	}

	return verificationToken, nil
}

func (s *UserService) GetUser(id string) (*schemas.UserResponse, error) {
	dbUser, err := s.repo.GetByID(uuid.MustParse(id))
	if err != nil {
		return nil, err
	}

	return &schemas.UserResponse{
		ID:            dbUser.ID,
		Username:      dbUser.Username,
		Email:         dbUser.Email,
		UserRole:      dbUser.UserRole,
		EmailVerified: dbUser.EmailVerified,
	}, nil
}

func (s *UserService) GetUserByUsername(username string) (*schemas.UserResponse, error) {
	dbUser, err := s.repo.GetByName(username)
	if err != nil {
		return nil, err
	}

	return &schemas.UserResponse{
		ID:            dbUser.ID,
		Username:      dbUser.Username,
		Password:      dbUser.Password,
		Email:         dbUser.Email,
		UserRole:      dbUser.UserRole,
		EmailVerified: dbUser.EmailVerified,
	}, nil
}

func (s *UserService) GetUserByEmail(email string) (*schemas.UserResponse, error) {
	dbUser, err := s.repo.GetByEmail(email)
	if err != nil {
		return nil, err
	}

	return &schemas.UserResponse{
		ID:            dbUser.ID,
		Username:      dbUser.Username,
		Email:         dbUser.Email,
		UserRole:      dbUser.UserRole,
		EmailVerified: dbUser.EmailVerified,
	}, nil
}

func (s *UserService) UpdateUser(user *models.User) error {
	return s.repo.Update(user)
}

func (s *UserService) DeleteUser(id string) error {
	return s.repo.Delete(uuid.MustParse(id))
}

func (s *UserService) ListUsers(page, pageSize int) ([]schemas.UserResponse, int64, error) {
	dbUsers, total, err := s.repo.List(page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	var users []schemas.UserResponse
	users = make([]schemas.UserResponse, total, total)

	for i, dbUser := range dbUsers {
		users[i] = schemas.UserResponse{
			ID:            dbUser.ID,
			Username:      dbUser.Username,
			Email:         dbUser.Email,
			UserRole:      dbUser.UserRole,
			EmailVerified: dbUser.EmailVerified,
		}
	}

	return users, total, nil
}

func (s *UserService) ResetPassword(token, newPassword string) error {
	user, err := s.repo.GetByResetToken(token, time.Now())
	if err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	if err := s.repo.UpdatePassword(user.ID, string(hashedPassword)); err != nil {
		return err
	}
	return s.repo.ClearResetToken(user.ID)
}

func (s *UserService) RequestPasswordReset(email string) (string, error) {
	user, err := s.repo.GetByEmail(email)
	if err != nil {
		return "", err
	}

	token := make([]byte, 32)
	if _, err := rand.Read(token); err != nil {
		return "", err
	}
	resetToken := base64.URLEncoding.EncodeToString(token)

	expiry := time.Now().Add(1 * time.Hour)

	if err := s.repo.SetResetToken(user.ID, resetToken, expiry); err != nil {
		return "", err
	}

	return resetToken, nil
}
