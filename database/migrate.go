package database

import (
	"gorm.io/gorm"
	"project/domain"
)

func Migrate(db *gorm.DB) error {
	var err error

	err = db.SetupJoinTable(&domain.Customer{}, "Redemptions", &domain.Redemption{})
	err = db.SetupJoinTable(&domain.Customer{}, "Vouchers", &domain.Order{})
	err = db.AutoMigrate(
		&domain.Voucher{},
		&domain.FreeShippingVoucher{},
		&domain.DiscountVoucher{},
		&domain.Customer{},
	)

	return err
}
