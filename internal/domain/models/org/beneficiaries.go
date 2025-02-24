package models

import (
	"github.com/google/uuid"
)

type Beneficiary struct {
	ID uuid.UUID `gorm:"type:uuid;primary_key;" json:"id"`

	Award       string `gorm:"type:text" json:"award"`
	Scholarship string `gorm:"type:text" json:"scholarship"`
}

func (e *Beneficiary) TableName() string {
	return "beneficiaries"
}
