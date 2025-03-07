package service

import (
	"encoding/json"
	"fmt"
	"jna-manager/internal/config"

	"github.com/go-resty/resty/v2"
)

type PaystackService struct {
	secretKey string
	client    *resty.Client
}

func NewPaystackService(cfg *config.Config) *PaystackService {
	return &PaystackService{
		secretKey: cfg.PaystackSecretKey,
		client:    resty.New(),
	}
}

func (s *PaystackService) InitiateTransaction(email, url string, amount float64) (string, string, error) {
	resp, err := s.client.R().
		SetHeader("Authorization", "Bearer "+s.secretKey).
		SetBody(map[string]any{
			"email":        email,
			"amount":       int(amount * 100),
			"callback_url": url,
		}).
		Post("https://api.paystack.co/transaction/initialize")

	if err != nil {
		return "", "", err
	}

	var paystackResp map[string]any
	err = json.Unmarshal(resp.Body(), &paystackResp)
	if err != nil {
		return "", "", err
	}

	data, ok := paystackResp["data"].(map[string]any)
	if !ok {
		return "", "", fmt.Errorf("unexpected response format")
	}

	reference, ok := data["reference"].(string)
	if !ok {
		return "", "", fmt.Errorf("reference not found in response")
	}

	authURL, ok := data["authorization_url"].(string)
	if !ok {
		return "", "", fmt.Errorf("authorization_url not found in response")
	}

	return reference, authURL, nil
}

func (s *PaystackService) VerifyTransaction(reference string) (*PaystackVerificationResponse, error) {
	resp, err := s.client.R().
		SetHeader("Authorization", "Bearer "+s.secretKey).
		Get(fmt.Sprintf("https://api.paystack.co/transaction/verify/%s", reference))

	if err != nil {
		return nil, err
	}

	var verificationResp PaystackVerificationResponse
	err = json.Unmarshal(resp.Body(), &verificationResp)
	if err != nil {
		return nil, err
	}

	return &verificationResp, nil
}

type PaystackVerificationResponse struct {
	Status bool `json:"status"`
	Data   struct {
		Status string `json:"status"`
		// Add other fields as needed
	} `json:"data"`
}
