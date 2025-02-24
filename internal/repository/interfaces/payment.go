package interfaces

import (
	models "jna-manager/internal/domain/models/payments"
)

type PaymentRepository interface {
	Create(payment *models.Payment) error
	GetByID(id string) (*models.Payment, error)
	GetByReference(reference string) (*models.Payment, error)

	Delete(id string) error
	List(page, pageSize int) ([]models.Payment, int64, error)
}
