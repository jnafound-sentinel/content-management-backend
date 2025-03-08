package interfaces

import (
	models "jna-manager/internal/domain/models/payments"

	"github.com/google/uuid"
)

type PaymentRepository interface {
	Create(payment *models.Payment) error
	GetByID(id uuid.UUID) (*models.Payment, error)
	GetByReference(reference string) (*models.Payment, error)
	Update(payment *models.Payment) error

	Delete(id uuid.UUID) error
	List(page, pageSize int) ([]models.Payment, int64, error)
}
