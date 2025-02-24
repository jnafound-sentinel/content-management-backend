package interfaces

import (
	models "jna-manager/internal/domain/models/payments"

	"github.com/google/uuid"
)

type DonationRepository interface {
	Create(donation *models.Donation) error
	GetByID(id uuid.UUID) (*models.Donation, error)
	GetByTagName(reference string) (*models.Donation, error)

	Delete(id uuid.UUID) error
	List(page, pageSize int) ([]models.Donation, int64, error)
}
