package repository

import (
	"gorm.io/gorm"
	"math"
	"project/domain"
	"project/util"
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

func (r *VoucherRepository) All(page uint, limit uint, isActive string, area string, voucherType string) (int64, int, uint, uint, []domain.Voucher, error) {
	var count int64
	r.db.Model(&domain.Voucher{}).Count(&count)
	pages := int(math.Ceil(float64(count) / float64(limit)))

	var vouchers []domain.Voucher
	result := r.db.Scopes(util.Paginate(page, limit)).Find(&vouchers)
	return count, pages, page, limit, vouchers, result.Error
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
