package repository

import (
	"gorm.io/gorm"
	"project/domain"
)

type RedeemVoucherRepository struct {
	db *gorm.DB
}

func NewRedeemVoucherRepository(db *gorm.DB) *RedeemVoucherRepository {
	return &RedeemVoucherRepository{db}
}

func (r *RedeemVoucherRepository) Create(voucher domain.Voucher) error {
	//return r.db.Model(&domain.Customer{ID: 1}).
	//	Association("Redemptions").
	//	Append(&domain.Redemption{VoucherId: 1})

	//return r.db.Model(&domain.Customer{ID: 1}).
	//	Association("Redemptions").
	//	Append([]domain.Redemption{{VoucherId: 1}})

	customer := domain.Customer{ID: 1}
	customer.Redemptions = append(customer.Redemptions, voucher)
	return r.db.Save(&customer).Error
}
