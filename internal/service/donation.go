package service

import (
	payments "jna-manager/internal/domain/models/payments"
	"jna-manager/internal/repository/interfaces"

	"github.com/google/uuid"
)

type DonationService struct {
	repo interfaces.DonationRepository
}

func NewDonationService(repo interfaces.DonationRepository) *DonationService {
	return &DonationService{repo: repo}
}

func (s *DonationService) CreateDonation(donation *payments.Donation) error {
	return s.repo.Create(donation)
}

func (s *DonationService) GetDonation(id string) (*payments.Donation, error) {
	return s.repo.GetByID(uuid.MustParse(id))
}

func (s *DonationService) GetDonationByTagName(tagName string) (*payments.Donation, error) {
	return s.repo.GetByTagName(tagName)
}

func (s *DonationService) UpdateDonation(donation *payments.Donation) (error) {
	return s.repo.Update(donation)
}

func (s *DonationService) DeleteDonation(id string) error {
	return s.repo.Delete(uuid.MustParse(id))
}

func (s *DonationService) ListDonations(page, pageSize int) ([]payments.Donation, int64, error) {
	return s.repo.List(page, pageSize)
}
