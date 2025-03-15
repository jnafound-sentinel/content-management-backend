package schemas

import "github.com/google/uuid"

type CreateBeneficiaryRequest struct {
	RecipientName string `json:"name" binding:"required"`
	Image         string `json:"image" binding:"required"`
	ProgramType   string `json:"program_type" binding:"required"`
	ShortBio      string `json:"short_bio" binding:"required"`
	FullBio       string `json:"full_bio" binding:"required"`
	Quote         string `json:"quote" binding:"required"`
	Featured      bool   `json:"featured" binding:"required"`
}

type UpdateBeneficiaryRequest struct {
	RecipientName string `json:"name" binding:"required"`
	Image         string `json:"image" binding:"required"`
	ProgramType   string `json:"program_type" binding:"required"`
	ShortBio      string `json:"short_bio" binding:"required"`
	FullBio       string `json:"full_bio" binding:"required"`
	Quote         string `json:"quote" binding:"required"`
	Featured      bool   `json:"featured" binding:"required"`
}

type BeneficiaryResponse struct {
	ID            uuid.UUID `json:"id"`
	RecipientName string    `json:"name" binding:"required"`
	Image         string    `json:"image" binding:"required"`
	ProgramType   string    `json:"program_type" binding:"required"`
	ShortBio      string    `json:"short_bio" binding:"required"`
	FullBio       string    `json:"full_bio" binding:"required"`
	Quote         string    `json:"quote" binding:"required"`
	Featured      bool      `json:"featured" binding:"required"`
}
