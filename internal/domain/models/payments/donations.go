package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Donation struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	TagName     string    `gorm:"size:50" json:"tag_name"`
	Purpose     string    `gorm:"size:100;default:'foundation'" json:"purpose"`
	Amount      int64     `gorm:"type:bigint" json:"amount"`
	Description string    `gorm:"type:text" json:"description"`
	Payments    []Payment `gorm:"foreignKey:DonationID" json:"payments"`
}

func (b *Donation) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.New()
	return
}

func (e *Donation) TableName() string {
	return "donations"
}
