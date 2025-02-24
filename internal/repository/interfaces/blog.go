package interfaces

import (
	models "jna-manager/internal/domain/models/org"

	"github.com/google/uuid"
)

type BlogRepository interface {
	Create(blog *models.Blog) error
	GetByID(id uuid.UUID) (*models.Blog, error)

	Update(blog *models.Blog) error
	Delete(id uuid.UUID) error

	List(page, pageSize int) ([]models.Blog, int64, error)
}
