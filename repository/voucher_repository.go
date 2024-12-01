package repository

import (
	"gorm.io/gorm"
	"project/domain"
)

type VoucherRepository struct {
	db *gorm.DB
}

func NewVoucherRepository(db *gorm.DB) *VoucherRepository {
	return &VoucherRepository{db}
}

func (r *VoucherRepository) Create(voucher domain.Voucher) error {
	return r.db.Create(&voucher).Error
}

func (r *VoucherRepository) All() ([]domain.Voucher, error) {
	var vouchers []domain.Voucher
	result := r.db.Find(&vouchers)
	return vouchers, result.Error
}

func (r *VoucherRepository) Get(id uint) (domain.Voucher, error) {
	var voucher domain.Voucher
	err := r.db.First(&voucher, id).Error
	return voucher, err
}

func (r *VoucherRepository) Update(voucher domain.Voucher) error {
	return r.db.Save(&voucher).Error
}

func (r *VoucherRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Voucher{}, id).Error
}
