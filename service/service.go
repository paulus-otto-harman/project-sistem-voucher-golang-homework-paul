package service

import "project/repository"

type Service struct {
	Voucher VoucherService
}

func NewService(repo repository.Repository) Service {
	return Service{
		Voucher: NewVoucherService(repo.Voucher),
	}
}
