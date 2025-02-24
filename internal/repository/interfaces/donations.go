package interfaces

import (
	models "jna-manager/internal/domain/models/payments"
)

type DonationRepository interface {
	Create(donation *models.Donation) error
	GetByID(id string) (*models.Donation, error)
	GetByTagName(reference string) (*models.Donation, error)

	Delete(id string) error
	List(page, pageSize int) ([]models.Donation, int64, error)
}
