package service

import (
	"bytes"
	"fmt"
	"html/template"
	"jna-manager/internal/config"
	"path/filepath"

	"github.com/resend/resend-go/v2"
)

type ResendClient struct {
	Client *resend.Client
	Domain string
}

type EmailService struct {
	client    ResendClient
	templates *template.Template
}

type ResetEmailData struct {
	ResetLink   string
	Username    string
	ExpiryHours int
}

func NewEmailService(cfg *config.Config) (*EmailService, error) {
	templates, err := template.ParseGlob(filepath.Join("templates", "emails", "*.html"))
	if err != nil {
		return nil, err
	}

	resendClient := ResendClient{
		Client: resend.NewClient(cfg.ResendApiKey),
		Domain: cfg.DomainID,
	}

	return &EmailService{
		client:    resendClient,
		templates: templates,
	}, nil
}

func (s *EmailService) SendPasswordResetEmail(to, resetToken, username string) (string, error) {
	resendDomain, err := s.client.Client.Domains.Get(s.client.Domain)
	if err != nil {
		return "", err
	}

	var body bytes.Buffer
	data := ResetEmailData{
		ResetLink:   fmt.Sprintf("https://%s/reset-password?token=%s", resendDomain.Name, resetToken),
		Username:    username,
		ExpiryHours: 1,
	}

	if err := s.templates.ExecuteTemplate(&body, "password-reset.html", data); err != nil {
		return "", err
	}

	params := &resend.SendEmailRequest{
		From:    fmt.Sprintf("Robolabs <info@%s>", resendDomain.Name),
		To:      []string{to},
		Html:    body.String(),
		Subject: "Reset Password",
	}

	sent, err := s.client.Client.Emails.Send(params)
	if err != nil {
		return "", err
	}

	return sent.Id, nil
}

func (s *EmailService) SendVerificationEmail(to, username, verificationToken string) (string, error) {
	resendDomain, err := s.client.Client.Domains.Get(s.client.Domain)
	if err != nil {
		return "", err
	}

	data := struct {
		Username         string
		VerificationLink string
		ExpiryHours      int
	}{
		Username:         username,
		VerificationLink: fmt.Sprintf("https://%s/verify-email?token=%s", resendDomain.Name, verificationToken),
		ExpiryHours:      24,
	}

	var body bytes.Buffer
	if err := s.templates.ExecuteTemplate(&body, "email-verification.html", data); err != nil {
		return "", err
	}

	params := &resend.SendEmailRequest{
		From:    fmt.Sprintf("Robolabs <info@%s>", resendDomain.Name),
		To:      []string{to},
		Html:    body.String(),
		Subject: "Reset Password",
	}

	sent, err := s.client.Client.Emails.Send(params)
	if err != nil {
		return "", err
	}

	return sent.Id, nil
}
