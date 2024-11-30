package database

import (
	"gorm.io/gorm"
	"project/domain"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&domain.Voucher{},
		&domain.FreeShippingVoucher{},
		&domain.DiscountVoucher{},
		&domain.Customer{},
		&domain.Redemption{},
	)
	return err
}
