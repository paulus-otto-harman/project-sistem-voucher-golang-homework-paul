package domain

import (
	"gorm.io/gorm"
	"time"
)

type Voucher struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	Name         string `gorm:"size:15;unique" json:"name"`
	Code         string `gorm:"size:15;unique" json:"code"`
	IsEcommerce  bool   `json:"type"`
	FreeShipping FreeShippingVoucher
	Discount     DiscountVoucher
	Area         string    `gorm:"size:15" json:"area"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	DeletedAt    gorm.DeletedAt
}

type Condition struct {
	IsEither      bool   `json:"is_either"`
	MinPurchase   uint32 `gorm:"precision:4" json:"min_purchase"`
	PaymentMethod string `gorm:"size:15" json:"payment_method"`
}

func VoucherSeed() []Voucher {
	return []Voucher{}
}
