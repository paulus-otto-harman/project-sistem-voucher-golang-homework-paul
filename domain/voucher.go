package domain

import (
	"gorm.io/gorm"
	"time"
)

type Voucher struct {
	ID           uint                `gorm:"primaryKey;autoIncrement" json:"id"`
	Name         string              `gorm:"size:15;unique" json:"name"`
	Code         string              `gorm:"size:15;unique" json:"code"`
	IsEcommerce  bool                `json:"type"`
	FreeShipping FreeShippingVoucher `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Discount     DiscountVoucher     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Area         string              `gorm:"size:15" json:"area"`
	CreatedAt    time.Time           `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time           `json:"updated_at"`
	DeletedAt    gorm.DeletedAt
}

func VoucherSeed() []Voucher {
	return []Voucher{
		{
			Name: "freeongkir1", Code: "code1", IsEcommerce: true,
			FreeShipping: FreeShippingVoucher{Condition: Condition{IsEither: true, MinPurchase: 10, PaymentMethod: "qris"}},
		},
		{
			Name: "freeongkir2", Code: "code2", IsEcommerce: true,
			FreeShipping: FreeShippingVoucher{Condition: Condition{IsEither: false, MinPurchase: 10, PaymentMethod: "qris"}},
		},
		{
			Name: "freeongkir3", Code: "code3", IsEcommerce: true,
			FreeShipping: FreeShippingVoucher{Condition: Condition{IsEither: true, MinPurchase: 10, PaymentMethod: "qris"}},
		},
		{
			Name: "freeongkir4", Code: "code4", IsEcommerce: true,
			FreeShipping: FreeShippingVoucher{Condition: Condition{IsEither: true, MinPurchase: 10, PaymentMethod: "qris"}},
		},
		{
			Name: "freeongkir5", Code: "code5", IsEcommerce: true,
			FreeShipping: FreeShippingVoucher{Condition: Condition{IsEither: true, MinPurchase: 10, PaymentMethod: "qris"}},
		},
		{
			Name: "freeongkir6", Code: "code6", IsEcommerce: true,
			FreeShipping: FreeShippingVoucher{Condition: Condition{IsEither: true, MinPurchase: 10, PaymentMethod: "qris"}},
		},
		{
			Name: "freeongkir7", Code: "code7", IsEcommerce: true,
			FreeShipping: FreeShippingVoucher{Condition: Condition{IsEither: true, MinPurchase: 10, PaymentMethod: "qris"}},
		},
	}
}
