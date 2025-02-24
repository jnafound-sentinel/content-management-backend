package schemas

import "github.com/google/uuid"

type CreatePaymentRequest struct {
	DonationID uuid.UUID `json:"donation_id" binding:"required"`
	Amount     int64     `json:"amount" binding:"required"`
	Reference  string    `json:"reference" binding:"required"`
}

type PaymentResponse struct {
	ID         uuid.UUID `json:"id"`
	DonationID uuid.UUID `json:"donation_id"`
	Amount     int64     `json:"amount"`
	Reference  string    `json:"reference"`
	Status     string    `json:"status"`
}
