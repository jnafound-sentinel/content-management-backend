package interfaces

import (
	models "jna-manager/internal/domain/models/org"
)

type BeneficiaryRepository interface {
	Create(beneficiary *models.Beneficiary) error
	GetByID(id string) (*models.Beneficiary, error)

	Update(beneficiary *models.Beneficiary) error
	Delete(id string) error

	List(page, pageSize int) ([]models.Beneficiary, int64, error)
}
