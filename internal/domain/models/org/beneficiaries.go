package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Beneficiary struct {
	ID uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`

	Award       string `gorm:"type:text" json:"award"`
	Scholarship string `gorm:"type:text" json:"scholarship"`
}

func (b *Beneficiary) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.New()
	return
}

func (e *Beneficiary) TableName() string {
	return "beneficiaries"
}
