package interfaces

import (
	models "jna-manager/internal/domain/models/org"
)

type BlogRepository interface {
	Create(blog *models.Blog) error
	GetByID(id string) (*models.Blog, error)

	Update(blog *models.Blog) error
	Delete(id string) error

	List(page, pageSize int) ([]models.Blog, int64, error)
}
