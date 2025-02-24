package service

import (
	org "jna-manager/internal/domain/models/org"
	"jna-manager/internal/repository/interfaces"

	"github.com/google/uuid"
)

type BeneficiaryService struct {
	repo interfaces.BeneficiaryRepository
}

func NewBeneficiaryService(repo interfaces.BeneficiaryRepository) *BeneficiaryService {
	return &BeneficiaryService{repo: repo}
}

func (s *BeneficiaryService) CreateBeneficiary(beneficiary *org.Beneficiary) error {
	return s.repo.Create(beneficiary)
}

func (s *BeneficiaryService) GetBeneficiary(id string) (*org.Beneficiary, error) {
	return s.repo.GetByID(uuid.MustParse(id))
}

func (s *BeneficiaryService) UpdateBeneficiary(beneficiary *org.Beneficiary) error {
	return s.repo.Update(beneficiary)
}

func (s *BeneficiaryService) DeleteBeneficiary(id string) error {
	return s.repo.Delete(uuid.MustParse(id))
}

func (s *BeneficiaryService) ListBeneficiaries(page, pageSize int) ([]org.Beneficiary, int64, error) {
	return s.repo.List(page, pageSize)
}
