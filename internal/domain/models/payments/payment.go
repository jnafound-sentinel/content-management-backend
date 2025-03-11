package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Payment struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	DonationID uuid.UUID `gorm:"type:uuid;foreignKey:DonationID" json:"donation_id"`
	Amount     int64     `json:"amount"`
	Reference  string    `json:"reference" gorm:"uniqueIndex"`
	Status     string    `json:"status"`
	Donation   Donation  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"donation"`
}

func (b *Payment) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.New()
	return
}

func (e *Payment) TableName() string {
	return "payments"
}
