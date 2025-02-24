package database

import (
	orgModels "jna-manager/internal/domain/models/org"
	"jna-manager/internal/repository/interfaces"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type beneficiaryRepository struct {
	db *gorm.DB
}

func NewBeneficiaryRepository(db *gorm.DB) interfaces.BeneficiaryRepository {
	return &beneficiaryRepository{db: db}
}

func (r *beneficiaryRepository) Create(beneficiary *orgModels.Beneficiary) error {
	return r.db.Create(beneficiary).Error
}

func (r *beneficiaryRepository) GetByID(id uuid.UUID) (*orgModels.Beneficiary, error) {
	var beneficiary orgModels.Beneficiary
	err := r.db.First(&beneficiary, id).Error
	if err != nil {
		return nil, err
	}
	return &beneficiary, nil
}

func (r *beneficiaryRepository) Update(beneficiary *orgModels.Beneficiary) error {
	return r.db.Save(beneficiary).Error
}

func (r *beneficiaryRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&orgModels.Beneficiary{}, id).Error
}

func (r *beneficiaryRepository) List(page, pageSize int) ([]orgModels.Beneficiary, int64, error) {
	var beneficiaries []orgModels.Beneficiary
	var total int64

	err := r.db.Model(&orgModels.Beneficiary{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	err = r.db.Offset(offset).Limit(pageSize).Find(&beneficiaries).Error
	if err != nil {
		return nil, 0, err
	}

	return beneficiaries, total, nil
}
