package database

import (
	"errors"
	paymentModels "jna-manager/internal/domain/models/payments"
	"jna-manager/internal/repository/interfaces"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type donationRepository struct {
	db *gorm.DB
}

func NewDonationRepository(db *gorm.DB) interfaces.DonationRepository {
	return &donationRepository{db: db}
}

func (r *donationRepository) Create(donation *paymentModels.Donation) error {
	return r.db.Create(donation).Error
}

func (r *donationRepository) GetByID(id uuid.UUID) (*paymentModels.Donation, error) {
	var donation paymentModels.Donation
	err := r.db.Preload("Payments").First(&donation, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("donation not found")
		}
		return nil, err
	}
	return &donation, nil
}

func (r *donationRepository) GetByTagName(tagName string) (*paymentModels.Donation, error) {
	var donation paymentModels.Donation
	err := r.db.Where("tag_name = ?", tagName).First(&donation).Error
	if err != nil {
		return nil, err
	}
	return &donation, nil
}

func (r *donationRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&paymentModels.Donation{}, id).Error
}

func (r *donationRepository) List(page, pageSize int) ([]paymentModels.Donation, int64, error) {
	var donations []paymentModels.Donation
	var total int64

	err := r.db.Model(&paymentModels.Donation{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	err = r.db.Offset(offset).Limit(pageSize).Find(&donations).Error
	if err != nil {
		return nil, 0, err
	}

	return donations, total, nil
}
