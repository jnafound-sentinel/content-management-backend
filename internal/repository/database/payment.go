package database

import (
	"errors"
	paymentModels "jna-manager/internal/domain/models/payments"
	"jna-manager/internal/repository/interfaces"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type paymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) interfaces.PaymentRepository {
	return &paymentRepository{db: db}
}

func (r *paymentRepository) Create(payment *paymentModels.Payment) error {
	return r.db.Create(payment).Error
}

func (r *paymentRepository) GetByID(id uuid.UUID) (*paymentModels.Payment, error) {
	var payment paymentModels.Payment
	err := r.db.Preload("Donation").First(&payment, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("payment not found")
		}
		return nil, err
	}
	return &payment, nil
}

func (r *paymentRepository) Update(payment *paymentModels.Payment) error {
	return r.db.Save(payment).Error
}

func (r *paymentRepository) GetByReference(reference string) (*paymentModels.Payment, error) {
	var payment paymentModels.Payment
	err := r.db.Where("reference = ?", reference).First(&payment).Error
	if err != nil {
		return nil, err
	}
	return &payment, nil
}

func (r *paymentRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&paymentModels.Payment{}, id).Error
}

func (r *paymentRepository) List(page, pageSize int) ([]paymentModels.Payment, int64, error) {
	var payments []paymentModels.Payment
	var total int64

	err := r.db.Model(&paymentModels.Payment{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	err = r.db.Offset(offset).Limit(pageSize).Find(&payments).Error
	if err != nil {
		return nil, 0, err
	}

	return payments, total, nil
}
