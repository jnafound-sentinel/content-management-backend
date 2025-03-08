package service

import (
	payments "jna-manager/internal/domain/models/payments"
	"jna-manager/internal/repository/interfaces"

	"github.com/google/uuid"
)

type PaymentService struct {
	repo interfaces.PaymentRepository
}

func NewPaymentService(repo interfaces.PaymentRepository) *PaymentService {
	return &PaymentService{repo: repo}
}

func (s *PaymentService) CreatePayment(payment *payments.Payment) error {
	return s.repo.Create(payment)
}

func (s *PaymentService) GetPayment(id string) (*payments.Payment, error) {
	return s.repo.GetByID(uuid.MustParse(id))
}

func (s *PaymentService) UpdatePayment(payment *payments.Payment) (error) {
	return s.repo.Update(payment)
}

func (s *PaymentService) GetPaymentByReference(reference string) (*payments.Payment, error) {
	return s.repo.GetByReference(reference)
}

func (s *PaymentService) DeletePayment(id string) error {
	return s.repo.Delete(uuid.MustParse(id))
}

func (s *PaymentService) ListPayments(page, pageSize int) ([]payments.Payment, int64, error) {
	return s.repo.List(page, pageSize)
}
