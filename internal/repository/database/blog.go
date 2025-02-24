package database

import (
	"errors"
	orgModels "jna-manager/internal/domain/models/org"
	"jna-manager/internal/repository/interfaces"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type blogRepository struct {
	db *gorm.DB
}

func NewBlogRepository(db *gorm.DB) interfaces.BlogRepository {
	return &blogRepository{db: db}
}

func (r *blogRepository) Create(blog *orgModels.Blog) error {
	return r.db.Create(blog).Error
}

func (r *blogRepository) GetByID(id uuid.UUID) (*orgModels.Blog, error) {
	var blog orgModels.Blog
	err := r.db.First(&blog, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("blog not found")
		}
		return nil, err
	}
	return &blog, nil
}

func (r *blogRepository) Update(blog *orgModels.Blog) error {
	return r.db.Save(blog).Error
}

func (r *blogRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&orgModels.Blog{}, id).Error
}

func (r *blogRepository) List(page, pageSize int) ([]orgModels.Blog, int64, error) {
	var blogs []orgModels.Blog
	var total int64

	err := r.db.Model(&orgModels.Blog{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	err = r.db.Offset(offset).Limit(pageSize).Find(&blogs).Error
	if err != nil {
		return nil, 0, err
	}

	return blogs, total, nil
}
