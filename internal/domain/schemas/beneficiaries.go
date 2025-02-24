package schemas

import "github.com/google/uuid"

type CreateBeneficiaryRequest struct {
	Award       string `json:"award" binding:"required"`
	Scholarship string `json:"scholarship" binding:"required"`
}

type UpdateBeneficiaryRequest struct {
	Award       string `json:"award"`
	Scholarship string `json:"scholarship"`
}

type BeneficiaryResponse struct {
	ID          uuid.UUID `json:"id"`
	Award       string    `json:"award"`
	Scholarship string    `json:"scholarship"`
}
