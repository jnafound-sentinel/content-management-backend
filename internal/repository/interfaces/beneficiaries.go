package interfaces

import (
	models "jna-manager/internal/domain/models/org"

	"github.com/google/uuid"
)

type BeneficiaryRepository interface {
	Create(beneficiary *models.Beneficiary) error
	GetByID(id uuid.UUID) (*models.Beneficiary, error)

	Update(beneficiary *models.Beneficiary) error
	Delete(id uuid.UUID) error

	List(page, pageSize int) ([]models.Beneficiary, int64, error)
}
