package service

import "project/repository"

type Service struct {
	Voucher       VoucherService
	RedeemVoucher RedeemVoucherService
}

func NewService(repo repository.Repository) Service {
	return Service{
		Voucher:       NewVoucherService(repo.Voucher),
		RedeemVoucher: NewRedeemVoucherService(repo.RedeemVoucher),
	}
}
