package schemas

import "github.com/google/uuid"

type CreateDonationRequest struct {
	TagName     string `json:"tag_name" binding:"required"`
	Purpose     string `json:"purpose" binding:"required"`
	Amount      int64  `json:"amount" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type DonationResponse struct {
	ID          uuid.UUID `json:"id"`
	TagName     string    `json:"tag_name"`
	Purpose     string    `json:"purpose"`
	Amount      int64     `json:"amount"`
	Description string    `json:"description"`
}
