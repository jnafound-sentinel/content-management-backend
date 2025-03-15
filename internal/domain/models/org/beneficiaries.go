package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Beneficiary struct {
	ID            uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	RecipientName string    `gorm:"type:text" json:"name"`
	Image         string    `gorm:"type:text" json:"image"`
	ProgramType   string    `gorm:"type:text" json:"program_type"`
	ShortBio      string    `gorm:"type:text" json:"short_bio"`
	FullBio       string    `gorm:"type:text" json:"full_bio"`
	Quote         string    `gorm:"type:text" json:"quote"`
	Featured      bool      `gorm:"type:boolean;default:false" json:"featured"`
}

func (b *Beneficiary) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.New()
	return
}

func (e *Beneficiary) TableName() string {
	return "beneficiaries"
}
