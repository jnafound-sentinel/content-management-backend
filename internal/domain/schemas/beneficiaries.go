package schemas

import "github.com/google/uuid"

type CreateBeneficiaryRequest struct {
	RecipientName string `json:"recipient" binding:"required"`
	Image         string `json:"image" binding:"required"`
	Award         string `json:"award" binding:"required"`
	Scholarship   string `json:"scholarship" binding:"required"`
}

type UpdateBeneficiaryRequest struct {
	RecipientName string `json:"recipient" binding:"required"`
	Image         string `json:"image" binding:"required"`
	Award         string `json:"award"`
	Scholarship   string `json:"scholarship"`
}

type BeneficiaryResponse struct {
	ID            uuid.UUID `json:"id"`
	RecipientName string    `json:"recipient" binding:"required"`
	Image         string    `json:"image" binding:"required"`
	Award         string    `json:"award"`
	Scholarship   string    `json:"scholarship"`
}
