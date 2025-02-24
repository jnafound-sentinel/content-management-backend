package models

import (
	"github.com/google/uuid"
)

type Donation struct {
	ID uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`

	TagName string `gorm:"size:50" json:"tag_name"`
	Purpose string `gorm:"size:100;default:'foundation'" json:"purpose"`

	Amount      int64  `gorm:"type:bigint" json:"amount"`
	Status      string `gorm:"size:100" json:"status"`
	Description string `gorm:"type:text" json:"description"`

	Payments []Payment `gorm:"foreignKey:DonationID" json:"payments"`
}

func (e *Donation) TableName() string {
	return "donations"
}
