package service

import (
	"project/domain"
	"project/repository"
)

type RedeemVoucherService interface {
	Create(voucher domain.Voucher) error
}

type redeemVoucherService struct {
	repo repository.RedeemVoucherRepository
}

func NewRedeemVoucherService(repo repository.RedeemVoucherRepository) RedeemVoucherService {
	return &redeemVoucherService{repo}
}

func (s *redeemVoucherService) Create(voucher domain.Voucher) error {
	return s.repo.Create(voucher)
}
