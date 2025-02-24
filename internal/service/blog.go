package service

import (
	org "jna-manager/internal/domain/models/org"
	"jna-manager/internal/repository/interfaces"

	"github.com/google/uuid"
)

type BlogService struct {
	repo interfaces.BlogRepository
}

func NewBlogService(repo interfaces.BlogRepository) *BlogService {
	return &BlogService{repo: repo}
}

func (s *BlogService) CreateBlog(blog *org.Blog) error {
	return s.repo.Create(blog)
}

func (s *BlogService) GetBlog(id string) (*org.Blog, error) {
	return s.repo.GetByID(uuid.MustParse(id))
}

func (s *BlogService) UpdateBlog(blog *org.Blog) error {
	return s.repo.Update(blog)
}

func (s *BlogService) DeleteBlog(id string) error {
	return s.repo.Delete(uuid.MustParse(id))
}

func (s *BlogService) ListBlogs(page, pageSize int) ([]org.Blog, int64, error) {
	return s.repo.List(page, pageSize)
}
