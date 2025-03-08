package schemas

type CreatePaymentRequest struct {
	DonationID  string `json:"donation_id" binding:"required"`
	Email       string `json:"email" binding:"required"`
	CallbackUrl string `json:"callback_url" binding:"required"`
	Amount      int64  `json:"amount" binding:"required"`
	Reference   string `json:"reference" binding:"required"`
}

type PaymentResponse struct {
	DonationID string `json:"donation_id"`
	Amount     int64  `json:"amount"`
	Reference  string `json:"reference"`
	AuthURL    string `json:"auth_url"`
	Status     string `json:"status"`
}

type VerifyPaymentRequest struct {
	Reference string `json:"reference"`
	PaymentID string `json:"payment_id"`
}
