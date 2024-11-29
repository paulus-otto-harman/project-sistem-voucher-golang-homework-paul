package domain

import (
	"gorm.io/gorm"
	"project/util"
	"time"
)

type Voucher struct {
	ID           uint                `gorm:"primaryKey;autoIncrement" json:"id"`
	Name         string              `gorm:"size:15;unique" json:"name"`
	Code         string              `gorm:"size:15;unique" json:"code"`
	IsEcommerce  bool                `json:"type"`
	FreeShipping FreeShippingVoucher `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Discount     DiscountVoucher     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	StartsAt     time.Time
	ExpiresAt    time.Time
	Area         string    `gorm:"size:15" json:"area"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	DeletedAt    gorm.DeletedAt
}

func VoucherSeed() []Voucher {
	return []Voucher{
		{
			Name: "afooj10q", Code: "code001", IsEcommerce: true,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			FreeShipping: FreeShippingVoucher{Condition: Condition{IsEither: true, MinPurchase: 10, PaymentMethod: "qris"}},
		},
		{
			Name: "afoaj10q", Code: "code002", IsEcommerce: true,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			FreeShipping: FreeShippingVoucher{Condition: Condition{IsEither: false, MinPurchase: 10, PaymentMethod: "qris"}},
		},
		{
			Name: "afooj25q", Code: "code003", IsEcommerce: true,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			FreeShipping: FreeShippingVoucher{Condition: Condition{IsEither: true, MinPurchase: 25, PaymentMethod: "qris"}},
		},
		{
			Name: "afoaj25q", Code: "code004", IsEcommerce: false,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			FreeShipping: FreeShippingVoucher{Condition: Condition{IsEither: false, MinPurchase: 25, PaymentMethod: "qris"}},
		},
		{
			Name: "afooj10c", Code: "code005", IsEcommerce: true,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			FreeShipping: FreeShippingVoucher{Condition: Condition{IsEither: true, MinPurchase: 10, PaymentMethod: "cod"}},
		},
		{
			Name: "afoaj10c", Code: "code006", IsEcommerce: true,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			FreeShipping: FreeShippingVoucher{Condition: Condition{IsEither: false, MinPurchase: 10, PaymentMethod: "cod"}},
		},
		{
			Name: "afooj25c", Code: "code007", IsEcommerce: true,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			FreeShipping: FreeShippingVoucher{Condition: Condition{IsEither: true, MinPurchase: 25, PaymentMethod: "cod"}},
		},
		{
			Name: "afoaj25c", Code: "code008", IsEcommerce: true,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			FreeShipping: FreeShippingVoucher{Condition: Condition{IsEither: false, MinPurchase: 25, PaymentMethod: "cod"}},
		},
		{
			Name: "afoojt10q", Code: "code009", IsEcommerce: true,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			FreeShipping: FreeShippingVoucher{Condition: Condition{IsEither: true, MinPurchase: 10, PaymentMethod: "qris"}},
		},
		{
			Name: "afoajt10q", Code: "code010", IsEcommerce: true,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			FreeShipping: FreeShippingVoucher{Condition: Condition{IsEither: false, MinPurchase: 10, PaymentMethod: "qris"}},
		},
		{
			Name: "afoojt25q", Code: "code011", IsEcommerce: true,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			FreeShipping: FreeShippingVoucher{Condition: Condition{IsEither: true, MinPurchase: 25, PaymentMethod: "qris"}},
		},
		{
			Name: "afoajt25q", Code: "code012", IsEcommerce: false,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			FreeShipping: FreeShippingVoucher{Condition: Condition{IsEither: false, MinPurchase: 25, PaymentMethod: "qris"}},
		},
		{
			Name: "afoojt10c", Code: "code013", IsEcommerce: true,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			FreeShipping: FreeShippingVoucher{Condition: Condition{IsEither: true, MinPurchase: 10, PaymentMethod: "cod"}},
		},
		{
			Name: "afoajt10c", Code: "code014", IsEcommerce: true,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			FreeShipping: FreeShippingVoucher{Condition: Condition{IsEither: false, MinPurchase: 10, PaymentMethod: "cod"}},
		},
		{
			Name: "afoojt25c", Code: "code015", IsEcommerce: true,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			FreeShipping: FreeShippingVoucher{Condition: Condition{IsEither: true, MinPurchase: 25, PaymentMethod: "cod"}},
		},
		{
			Name: "afoajt25c", Code: "code016", IsEcommerce: true,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			FreeShipping: FreeShippingVoucher{Condition: Condition{IsEither: false, MinPurchase: 25, PaymentMethod: "cod"}},
		},
		// //
		{
			Name: "xfooj10q", Code: "code017", IsEcommerce: true,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			FreeShipping: FreeShippingVoucher{Condition: Condition{IsEither: true, MinPurchase: 10, PaymentMethod: "qris"}},
		},
		{
			Name: "xfoaj10q", Code: "code018", IsEcommerce: true,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			FreeShipping: FreeShippingVoucher{Condition: Condition{IsEither: false, MinPurchase: 10, PaymentMethod: "qris"}},
		},
		{
			Name: "xfooj25q", Code: "code019", IsEcommerce: true,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			FreeShipping: FreeShippingVoucher{Condition: Condition{IsEither: true, MinPurchase: 25, PaymentMethod: "qris"}},
		},
		{
			Name: "xfoaj25q", Code: "code020", IsEcommerce: false,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			FreeShipping: FreeShippingVoucher{Condition: Condition{IsEither: false, MinPurchase: 25, PaymentMethod: "qris"}},
		},
		{
			Name: "xfooj10c", Code: "code021", IsEcommerce: true,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			FreeShipping: FreeShippingVoucher{Condition: Condition{IsEither: true, MinPurchase: 10, PaymentMethod: "cod"}},
		},
		{
			Name: "xfoaj10c", Code: "code022", IsEcommerce: true,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			FreeShipping: FreeShippingVoucher{Condition: Condition{IsEither: false, MinPurchase: 10, PaymentMethod: "cod"}},
		},
		{
			Name: "xfooj25c", Code: "code023", IsEcommerce: true,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			FreeShipping: FreeShippingVoucher{Condition: Condition{IsEither: true, MinPurchase: 25, PaymentMethod: "cod"}},
		},
		{
			Name: "xfoaj25c", Code: "code024", IsEcommerce: true,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			FreeShipping: FreeShippingVoucher{Condition: Condition{IsEither: false, MinPurchase: 25, PaymentMethod: "cod"}},
		},
		{
			Name: "xfoojt10q", Code: "code025", IsEcommerce: true,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			FreeShipping: FreeShippingVoucher{Condition: Condition{IsEither: true, MinPurchase: 10, PaymentMethod: "qris"}},
		},
		{
			Name: "xfoajt10q", Code: "code026", IsEcommerce: true,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			FreeShipping: FreeShippingVoucher{Condition: Condition{IsEither: false, MinPurchase: 10, PaymentMethod: "qris"}},
		},
		{
			Name: "xfoojt25q", Code: "code027", IsEcommerce: true,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			FreeShipping: FreeShippingVoucher{Condition: Condition{IsEither: true, MinPurchase: 25, PaymentMethod: "qris"}},
		},
		{
			Name: "xfoajt25q", Code: "code028", IsEcommerce: false,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			FreeShipping: FreeShippingVoucher{Condition: Condition{IsEither: false, MinPurchase: 25, PaymentMethod: "qris"}},
		},
		{
			Name: "xfoojt10c", Code: "code029", IsEcommerce: true,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			FreeShipping: FreeShippingVoucher{Condition: Condition{IsEither: true, MinPurchase: 10, PaymentMethod: "cod"}},
		},
		{
			Name: "xfoajt10c", Code: "code030", IsEcommerce: true,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			FreeShipping: FreeShippingVoucher{Condition: Condition{IsEither: false, MinPurchase: 10, PaymentMethod: "cod"}},
		},
		{
			Name: "xfoojt25c", Code: "code031", IsEcommerce: true,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			FreeShipping: FreeShippingVoucher{Condition: Condition{IsEither: true, MinPurchase: 25, PaymentMethod: "cod"}},
		},
		{
			Name: "xfoajt25c", Code: "code032", IsEcommerce: true,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			FreeShipping: FreeShippingVoucher{Condition: Condition{IsEither: false, MinPurchase: 25, PaymentMethod: "cod"}},
		},
		// // //
		{
			Name: "adro10j10q", Code: "code033", IsEcommerce: true,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: DiscountVoucher{Rate: 10, Condition: Condition{IsEither: true, MinPurchase: 10, PaymentMethod: "qris"}},
		},
		{
			Name: "adra10j10q", Code: "code034", IsEcommerce: true,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: DiscountVoucher{Rate: 10, Condition: Condition{IsEither: false, MinPurchase: 10, PaymentMethod: "qris"}},
		},
		{
			Name: "adro10j25q", Code: "code035", IsEcommerce: true,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: DiscountVoucher{Rate: 10, Condition: Condition{IsEither: true, MinPurchase: 25, PaymentMethod: "qris"}},
		},
		{
			Name: "adra10j25q", Code: "code036", IsEcommerce: false,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: DiscountVoucher{Rate: 10, Condition: Condition{IsEither: false, MinPurchase: 25, PaymentMethod: "qris"}},
		},
		{
			Name: "adro10j10c", Code: "code037", IsEcommerce: true,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: DiscountVoucher{Rate: 10, Condition: Condition{IsEither: true, MinPurchase: 10, PaymentMethod: "cod"}},
		},
		{
			Name: "adra10j10c", Code: "code038", IsEcommerce: true,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: DiscountVoucher{Rate: 10, Condition: Condition{IsEither: false, MinPurchase: 10, PaymentMethod: "cod"}},
		},
		{
			Name: "adro10j25c", Code: "code039", IsEcommerce: true,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: DiscountVoucher{Rate: 10, Condition: Condition{IsEither: true, MinPurchase: 25, PaymentMethod: "cod"}},
		},
		{
			Name: "adra10j25c", Code: "code040", IsEcommerce: true,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: DiscountVoucher{Rate: 10, Condition: Condition{IsEither: false, MinPurchase: 25, PaymentMethod: "cod"}},
		},
		{
			Name: "adro10jt10q", Code: "code041", IsEcommerce: true,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: DiscountVoucher{Rate: 10, Condition: Condition{IsEither: true, MinPurchase: 10, PaymentMethod: "qris"}},
		},
		{
			Name: "adra10jt10q", Code: "code042", IsEcommerce: true,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: DiscountVoucher{Rate: 10, Condition: Condition{IsEither: false, MinPurchase: 10, PaymentMethod: "qris"}},
		},
		{
			Name: "adro10jt25q", Code: "code043", IsEcommerce: true,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: DiscountVoucher{Rate: 10, Condition: Condition{IsEither: true, MinPurchase: 25, PaymentMethod: "qris"}},
		},
		{
			Name: "adra10jt25q", Code: "code044", IsEcommerce: false,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: DiscountVoucher{Rate: 10, Condition: Condition{IsEither: false, MinPurchase: 25, PaymentMethod: "qris"}},
		},
		{
			Name: "adro10jt10c", Code: "code045", IsEcommerce: true,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: DiscountVoucher{Rate: 10, Condition: Condition{IsEither: true, MinPurchase: 10, PaymentMethod: "cod"}},
		},
		{
			Name: "adra10jt10c", Code: "code046", IsEcommerce: true,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: DiscountVoucher{Rate: 10, Condition: Condition{IsEither: false, MinPurchase: 10, PaymentMethod: "cod"}},
		},
		{
			Name: "adro10jt25c", Code: "code047", IsEcommerce: true,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: DiscountVoucher{Rate: 10, Condition: Condition{IsEither: true, MinPurchase: 25, PaymentMethod: "cod"}},
		},
		{
			Name: "adra10jt25c", Code: "code048", IsEcommerce: true,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: DiscountVoucher{Rate: 10, Condition: Condition{IsEither: false, MinPurchase: 25, PaymentMethod: "cod"}},
		},
		{
			Name: "xdro10j10q", Code: "code049", IsEcommerce: true,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: DiscountVoucher{Rate: 10, Condition: Condition{IsEither: true, MinPurchase: 10, PaymentMethod: "qris"}},
		},
		{
			Name: "xdra10j10q", Code: "code050", IsEcommerce: true,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: DiscountVoucher{Rate: 10, Condition: Condition{IsEither: false, MinPurchase: 10, PaymentMethod: "qris"}},
		},
		{
			Name: "xdro10j25q", Code: "code051", IsEcommerce: true,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: DiscountVoucher{Rate: 10, Condition: Condition{IsEither: true, MinPurchase: 25, PaymentMethod: "qris"}},
		},
		{
			Name: "xdra10j25q", Code: "code052", IsEcommerce: false,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: DiscountVoucher{Rate: 10, Condition: Condition{IsEither: false, MinPurchase: 25, PaymentMethod: "qris"}},
		},
		{
			Name: "xdro10j10c", Code: "code053", IsEcommerce: true,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: DiscountVoucher{Rate: 10, Condition: Condition{IsEither: true, MinPurchase: 10, PaymentMethod: "cod"}},
		},
		{
			Name: "xdra10j10c", Code: "code054", IsEcommerce: true,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: DiscountVoucher{Rate: 10, Condition: Condition{IsEither: false, MinPurchase: 10, PaymentMethod: "cod"}},
		},
		{
			Name: "xdro10j25c", Code: "code055", IsEcommerce: true,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: DiscountVoucher{Rate: 10, Condition: Condition{IsEither: true, MinPurchase: 25, PaymentMethod: "cod"}},
		},
		{
			Name: "xdra10j25c", Code: "code056", IsEcommerce: true,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: DiscountVoucher{Rate: 10, Condition: Condition{IsEither: false, MinPurchase: 25, PaymentMethod: "cod"}},
		},
		{
			Name: "xdro10jt10q", Code: "code057", IsEcommerce: true,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: DiscountVoucher{Rate: 10, Condition: Condition{IsEither: true, MinPurchase: 10, PaymentMethod: "qris"}},
		},
		{
			Name: "xdra10jt10q", Code: "code058", IsEcommerce: true,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: DiscountVoucher{Rate: 10, Condition: Condition{IsEither: false, MinPurchase: 10, PaymentMethod: "qris"}},
		},
		{
			Name: "xdro10jt25q", Code: "code059", IsEcommerce: true,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: DiscountVoucher{Rate: 10, Condition: Condition{IsEither: true, MinPurchase: 25, PaymentMethod: "qris"}},
		},
		{
			Name: "xdra10jt25q", Code: "code060", IsEcommerce: false,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: DiscountVoucher{Rate: 10, Condition: Condition{IsEither: false, MinPurchase: 25, PaymentMethod: "qris"}},
		},
		{
			Name: "xdro10jt10c", Code: "code061", IsEcommerce: true,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: DiscountVoucher{Rate: 10, Condition: Condition{IsEither: true, MinPurchase: 10, PaymentMethod: "cod"}},
		},
		{
			Name: "xdra10jt10c", Code: "code062", IsEcommerce: true,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: DiscountVoucher{Rate: 10, Condition: Condition{IsEither: false, MinPurchase: 10, PaymentMethod: "cod"}},
		},
		{
			Name: "xdro10jt25c", Code: "code063", IsEcommerce: true,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: DiscountVoucher{Rate: 10, Condition: Condition{IsEither: true, MinPurchase: 25, PaymentMethod: "cod"}},
		},
		{
			Name: "xdra10jt25c", Code: "code064", IsEcommerce: true,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: DiscountVoucher{Rate: 10, Condition: Condition{IsEither: false, MinPurchase: 25, PaymentMethod: "cod"}},
		},
		//// // // //
		{
			Name: "adro25j10q", Code: "code065", IsEcommerce: true,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: DiscountVoucher{Rate: 25, Condition: Condition{IsEither: true, MinPurchase: 10, PaymentMethod: "qris"}},
		},
		{
			Name: "adra25j10q", Code: "code066", IsEcommerce: true,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: DiscountVoucher{Rate: 25, Condition: Condition{IsEither: false, MinPurchase: 10, PaymentMethod: "qris"}},
		},
		{
			Name: "adro25j25q", Code: "code067", IsEcommerce: true,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: DiscountVoucher{Rate: 25, Condition: Condition{IsEither: true, MinPurchase: 25, PaymentMethod: "qris"}},
		},
		{
			Name: "adra25j25q", Code: "code068", IsEcommerce: false,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: DiscountVoucher{Rate: 25, Condition: Condition{IsEither: false, MinPurchase: 25, PaymentMethod: "qris"}},
		},
		{
			Name: "adro25j10c", Code: "code069", IsEcommerce: true,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: DiscountVoucher{Rate: 25, Condition: Condition{IsEither: true, MinPurchase: 10, PaymentMethod: "cod"}},
		},
		{
			Name: "adra25j10c", Code: "code070", IsEcommerce: true,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: DiscountVoucher{Rate: 25, Condition: Condition{IsEither: false, MinPurchase: 10, PaymentMethod: "cod"}},
		},
		{
			Name: "adro25j25c", Code: "code071", IsEcommerce: true,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: DiscountVoucher{Rate: 25, Condition: Condition{IsEither: true, MinPurchase: 25, PaymentMethod: "cod"}},
		},
		{
			Name: "adra25j25c", Code: "code072", IsEcommerce: true,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: DiscountVoucher{Rate: 25, Condition: Condition{IsEither: false, MinPurchase: 25, PaymentMethod: "cod"}},
		},
		{
			Name: "adro25jt10q", Code: "code073", IsEcommerce: true,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: DiscountVoucher{Rate: 25, Condition: Condition{IsEither: true, MinPurchase: 10, PaymentMethod: "qris"}},
		},
		{
			Name: "adra25jt10q", Code: "code074", IsEcommerce: true,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: DiscountVoucher{Rate: 25, Condition: Condition{IsEither: false, MinPurchase: 10, PaymentMethod: "qris"}},
		},
		{
			Name: "adro25jt25q", Code: "code075", IsEcommerce: true,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: DiscountVoucher{Rate: 25, Condition: Condition{IsEither: true, MinPurchase: 25, PaymentMethod: "qris"}},
		},
		{
			Name: "adra25jt25q", Code: "code076", IsEcommerce: false,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: DiscountVoucher{Rate: 25, Condition: Condition{IsEither: false, MinPurchase: 25, PaymentMethod: "qris"}},
		},
		{
			Name: "adro25jt10c", Code: "code077", IsEcommerce: true,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: DiscountVoucher{Rate: 25, Condition: Condition{IsEither: true, MinPurchase: 10, PaymentMethod: "cod"}},
		},
		{
			Name: "adra25jt10c", Code: "code078", IsEcommerce: true,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: DiscountVoucher{Rate: 25, Condition: Condition{IsEither: false, MinPurchase: 10, PaymentMethod: "cod"}},
		},
		{
			Name: "adro25jt25c", Code: "code079", IsEcommerce: true,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: DiscountVoucher{Rate: 25, Condition: Condition{IsEither: true, MinPurchase: 25, PaymentMethod: "cod"}},
		},
		{
			Name: "adra25jt25c", Code: "code080", IsEcommerce: true,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: DiscountVoucher{Rate: 25, Condition: Condition{IsEither: false, MinPurchase: 25, PaymentMethod: "cod"}},
		},
		{
			Name: "xdro25j10q", Code: "code081", IsEcommerce: true,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: DiscountVoucher{Rate: 25, Condition: Condition{IsEither: true, MinPurchase: 10, PaymentMethod: "qris"}},
		},
		{
			Name: "xdra25j10q", Code: "code082", IsEcommerce: true,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: DiscountVoucher{Rate: 25, Condition: Condition{IsEither: false, MinPurchase: 10, PaymentMethod: "qris"}},
		},
		{
			Name: "xdro25j25q", Code: "code083", IsEcommerce: true,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: DiscountVoucher{Rate: 25, Condition: Condition{IsEither: true, MinPurchase: 25, PaymentMethod: "qris"}},
		},
		{
			Name: "xdra25j25q", Code: "code084", IsEcommerce: false,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: DiscountVoucher{Rate: 25, Condition: Condition{IsEither: false, MinPurchase: 25, PaymentMethod: "qris"}},
		},
		{
			Name: "xdro25j10c", Code: "code085", IsEcommerce: true,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: DiscountVoucher{Rate: 25, Condition: Condition{IsEither: true, MinPurchase: 10, PaymentMethod: "cod"}},
		},
		{
			Name: "xdra25j10c", Code: "code086", IsEcommerce: true,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: DiscountVoucher{Rate: 25, Condition: Condition{IsEither: false, MinPurchase: 10, PaymentMethod: "cod"}},
		},
		{
			Name: "xdro25j25c", Code: "code087", IsEcommerce: true,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: DiscountVoucher{Rate: 25, Condition: Condition{IsEither: true, MinPurchase: 25, PaymentMethod: "cod"}},
		},
		{
			Name: "xdra25j25c", Code: "code088", IsEcommerce: true,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: DiscountVoucher{Rate: 25, Condition: Condition{IsEither: false, MinPurchase: 25, PaymentMethod: "cod"}},
		},
		{
			Name: "xdro25jt10q", Code: "code089", IsEcommerce: true,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: DiscountVoucher{Rate: 25, Condition: Condition{IsEither: true, MinPurchase: 10, PaymentMethod: "qris"}},
		},
		{
			Name: "xdra25jt10q", Code: "code090", IsEcommerce: true,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: DiscountVoucher{Rate: 25, Condition: Condition{IsEither: false, MinPurchase: 10, PaymentMethod: "qris"}},
		},
		{
			Name: "xdro25jt25q", Code: "code091", IsEcommerce: true,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: DiscountVoucher{Rate: 25, Condition: Condition{IsEither: true, MinPurchase: 25, PaymentMethod: "qris"}},
		},
		{
			Name: "xdra25jt25q", Code: "code092", IsEcommerce: false,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: DiscountVoucher{Rate: 25, Condition: Condition{IsEither: false, MinPurchase: 25, PaymentMethod: "qris"}},
		},
		{
			Name: "xdro25jt10c", Code: "code093", IsEcommerce: true,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: DiscountVoucher{Rate: 25, Condition: Condition{IsEither: true, MinPurchase: 10, PaymentMethod: "cod"}},
		},
		{
			Name: "xdra25jt10c", Code: "code094", IsEcommerce: true,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: DiscountVoucher{Rate: 25, Condition: Condition{IsEither: false, MinPurchase: 10, PaymentMethod: "cod"}},
		},
		{
			Name: "xdro25jt25c", Code: "code095", IsEcommerce: true,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: DiscountVoucher{Rate: 25, Condition: Condition{IsEither: true, MinPurchase: 25, PaymentMethod: "cod"}},
		},
		{
			Name: "xdra25jt25c", Code: "code096", IsEcommerce: true,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: DiscountVoucher{Rate: 25, Condition: Condition{IsEither: false, MinPurchase: 25, PaymentMethod: "cod"}},
		},
		//// // // // //
		{
			Name: "adao10j100q", Code: "code097", IsEcommerce: true,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: DiscountVoucher{Amount: 10, Condition: Condition{IsEither: true, MinPurchase: 100, PaymentMethod: "qris"}},
		},
		{
			Name: "adaa10j100q", Code: "code098", IsEcommerce: true,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: DiscountVoucher{Amount: 10, Condition: Condition{IsEither: false, MinPurchase: 100, PaymentMethod: "qris"}},
		},
		{
			Name: "adao10j250q", Code: "code099", IsEcommerce: true,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: DiscountVoucher{Amount: 10, Condition: Condition{IsEither: true, MinPurchase: 250, PaymentMethod: "qris"}},
		},
		{
			Name: "adaa10j250q", Code: "code100", IsEcommerce: false,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: DiscountVoucher{Amount: 10, Condition: Condition{IsEither: false, MinPurchase: 250, PaymentMethod: "qris"}},
		},
		{
			Name: "adao10j100c", Code: "code101", IsEcommerce: true,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: DiscountVoucher{Amount: 10, Condition: Condition{IsEither: true, MinPurchase: 100, PaymentMethod: "cod"}},
		},
		{
			Name: "adaa10j100c", Code: "code102", IsEcommerce: true,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: DiscountVoucher{Amount: 10, Condition: Condition{IsEither: false, MinPurchase: 100, PaymentMethod: "cod"}},
		},
		{
			Name: "adao10j250c", Code: "code103", IsEcommerce: true,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: DiscountVoucher{Amount: 10, Condition: Condition{IsEither: true, MinPurchase: 250, PaymentMethod: "cod"}},
		},
		{
			Name: "adaa10j250c", Code: "code104", IsEcommerce: true,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: DiscountVoucher{Amount: 10, Condition: Condition{IsEither: false, MinPurchase: 250, PaymentMethod: "cod"}},
		},
		{
			Name: "adao10jt100q", Code: "code105", IsEcommerce: true,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: DiscountVoucher{Amount: 10, Condition: Condition{IsEither: true, MinPurchase: 100, PaymentMethod: "qris"}},
		},
		{
			Name: "adaa10jt100q", Code: "code106", IsEcommerce: true,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: DiscountVoucher{Amount: 10, Condition: Condition{IsEither: false, MinPurchase: 100, PaymentMethod: "qris"}},
		},
		{
			Name: "adao10jt250q", Code: "code107", IsEcommerce: true,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: DiscountVoucher{Amount: 10, Condition: Condition{IsEither: true, MinPurchase: 250, PaymentMethod: "qris"}},
		},
		{
			Name: "adaa10jt250q", Code: "code108", IsEcommerce: false,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: DiscountVoucher{Amount: 10, Condition: Condition{IsEither: false, MinPurchase: 250, PaymentMethod: "qris"}},
		},
		{
			Name: "adao10jt100c", Code: "code109", IsEcommerce: true,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: DiscountVoucher{Amount: 10, Condition: Condition{IsEither: true, MinPurchase: 100, PaymentMethod: "cod"}},
		},
		{
			Name: "adaa10jt100c", Code: "code110", IsEcommerce: true,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: DiscountVoucher{Amount: 10, Condition: Condition{IsEither: false, MinPurchase: 100, PaymentMethod: "cod"}},
		},
		{
			Name: "adao10jt250c", Code: "code111", IsEcommerce: true,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: DiscountVoucher{Amount: 10, Condition: Condition{IsEither: true, MinPurchase: 250, PaymentMethod: "cod"}},
		},
		{
			Name: "adaa10jt250c", Code: "code112", IsEcommerce: true,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: DiscountVoucher{Amount: 10, Condition: Condition{IsEither: false, MinPurchase: 250, PaymentMethod: "cod"}},
		},
		{
			Name: "xdao10j100q", Code: "code113", IsEcommerce: true,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: DiscountVoucher{Amount: 10, Condition: Condition{IsEither: true, MinPurchase: 100, PaymentMethod: "qris"}},
		},
		{
			Name: "xdaa10j100q", Code: "code114", IsEcommerce: true,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: DiscountVoucher{Amount: 10, Condition: Condition{IsEither: false, MinPurchase: 100, PaymentMethod: "qris"}},
		},
		{
			Name: "xdao10j250q", Code: "code115", IsEcommerce: true,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: DiscountVoucher{Amount: 10, Condition: Condition{IsEither: true, MinPurchase: 250, PaymentMethod: "qris"}},
		},
		{
			Name: "xdaa10j250q", Code: "code116", IsEcommerce: false,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: DiscountVoucher{Amount: 10, Condition: Condition{IsEither: false, MinPurchase: 250, PaymentMethod: "qris"}},
		},
		{
			Name: "xdao10j100c", Code: "code117", IsEcommerce: true,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: DiscountVoucher{Amount: 10, Condition: Condition{IsEither: true, MinPurchase: 100, PaymentMethod: "cod"}},
		},
		{
			Name: "xdaa10j100c", Code: "code118", IsEcommerce: true,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: DiscountVoucher{Amount: 10, Condition: Condition{IsEither: false, MinPurchase: 100, PaymentMethod: "cod"}},
		},
		{
			Name: "xdao10j250c", Code: "code119", IsEcommerce: true,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: DiscountVoucher{Amount: 10, Condition: Condition{IsEither: true, MinPurchase: 250, PaymentMethod: "cod"}},
		},
		{
			Name: "xdaa10j250c", Code: "code120", IsEcommerce: true,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: DiscountVoucher{Amount: 10, Condition: Condition{IsEither: false, MinPurchase: 250, PaymentMethod: "cod"}},
		},
		{
			Name: "xdao10jt100q", Code: "code121", IsEcommerce: true,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: DiscountVoucher{Amount: 10, Condition: Condition{IsEither: true, MinPurchase: 100, PaymentMethod: "qris"}},
		},
		{
			Name: "xdaa10jt100q", Code: "code122", IsEcommerce: true,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: DiscountVoucher{Amount: 10, Condition: Condition{IsEither: false, MinPurchase: 100, PaymentMethod: "qris"}},
		},
		{
			Name: "xdao10jt250q", Code: "code123", IsEcommerce: true,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: DiscountVoucher{Amount: 10, Condition: Condition{IsEither: true, MinPurchase: 250, PaymentMethod: "qris"}},
		},
		{
			Name: "xdaa10jt250q", Code: "code124", IsEcommerce: false,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: DiscountVoucher{Amount: 10, Condition: Condition{IsEither: false, MinPurchase: 250, PaymentMethod: "qris"}},
		},
		{
			Name: "xdao10jt100c", Code: "code125", IsEcommerce: true,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: DiscountVoucher{Amount: 10, Condition: Condition{IsEither: true, MinPurchase: 100, PaymentMethod: "cod"}},
		},
		{
			Name: "xdaa10jt100c", Code: "code126", IsEcommerce: true,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: DiscountVoucher{Amount: 10, Condition: Condition{IsEither: false, MinPurchase: 100, PaymentMethod: "cod"}},
		},
		{
			Name: "xdao10jt250c", Code: "code127", IsEcommerce: true,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: DiscountVoucher{Amount: 10, Condition: Condition{IsEither: true, MinPurchase: 250, PaymentMethod: "cod"}},
		},
		{
			Name: "xdaa10jt250c", Code: "code128", IsEcommerce: true,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: DiscountVoucher{Amount: 10, Condition: Condition{IsEither: false, MinPurchase: 250, PaymentMethod: "cod"}},
		},
		//// // // // // //
		{
			Name: "adao25j200q", Code: "code129", IsEcommerce: true,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: DiscountVoucher{Amount: 25, Condition: Condition{IsEither: true, MinPurchase: 200, PaymentMethod: "qris"}},
		},
		{
			Name: "adaa25j200q", Code: "code130", IsEcommerce: true,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: DiscountVoucher{Amount: 25, Condition: Condition{IsEither: false, MinPurchase: 200, PaymentMethod: "qris"}},
		},
		{
			Name: "adao25j400q", Code: "code131", IsEcommerce: true,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: DiscountVoucher{Amount: 25, Condition: Condition{IsEither: true, MinPurchase: 400, PaymentMethod: "qris"}},
		},
		{
			Name: "adaa25j400q", Code: "code132", IsEcommerce: false,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: DiscountVoucher{Amount: 25, Condition: Condition{IsEither: false, MinPurchase: 400, PaymentMethod: "qris"}},
		},
		{
			Name: "adao25j200c", Code: "code133", IsEcommerce: true,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: DiscountVoucher{Amount: 25, Condition: Condition{IsEither: true, MinPurchase: 200, PaymentMethod: "cod"}},
		},
		{
			Name: "adaa25j200c", Code: "code134", IsEcommerce: true,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: DiscountVoucher{Amount: 25, Condition: Condition{IsEither: false, MinPurchase: 200, PaymentMethod: "cod"}},
		},
		{
			Name: "adao25j400c", Code: "code135", IsEcommerce: true,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: DiscountVoucher{Amount: 25, Condition: Condition{IsEither: true, MinPurchase: 400, PaymentMethod: "cod"}},
		},
		{
			Name: "adaa25j400c", Code: "code136", IsEcommerce: true,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: DiscountVoucher{Amount: 25, Condition: Condition{IsEither: false, MinPurchase: 400, PaymentMethod: "cod"}},
		},
		{
			Name: "adao25jt200q", Code: "code137", IsEcommerce: true,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: DiscountVoucher{Amount: 25, Condition: Condition{IsEither: true, MinPurchase: 200, PaymentMethod: "qris"}},
		},
		{
			Name: "adaa25jt200q", Code: "code138", IsEcommerce: true,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: DiscountVoucher{Amount: 25, Condition: Condition{IsEither: false, MinPurchase: 200, PaymentMethod: "qris"}},
		},
		{
			Name: "adao25jt400q", Code: "code139", IsEcommerce: true,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: DiscountVoucher{Amount: 25, Condition: Condition{IsEither: true, MinPurchase: 400, PaymentMethod: "qris"}},
		},
		{
			Name: "adaa25jt400q", Code: "code140", IsEcommerce: false,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: DiscountVoucher{Amount: 25, Condition: Condition{IsEither: false, MinPurchase: 400, PaymentMethod: "qris"}},
		},
		{
			Name: "adao25jt200c", Code: "code141", IsEcommerce: true,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: DiscountVoucher{Amount: 25, Condition: Condition{IsEither: true, MinPurchase: 200, PaymentMethod: "cod"}},
		},
		{
			Name: "adaa25jt200c", Code: "code142", IsEcommerce: true,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: DiscountVoucher{Amount: 25, Condition: Condition{IsEither: false, MinPurchase: 200, PaymentMethod: "cod"}},
		},
		{
			Name: "adao25jt400c", Code: "code143", IsEcommerce: true,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: DiscountVoucher{Amount: 25, Condition: Condition{IsEither: true, MinPurchase: 400, PaymentMethod: "cod"}},
		},
		{
			Name: "adaa25jt400c", Code: "code144", IsEcommerce: true,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: DiscountVoucher{Amount: 25, Condition: Condition{IsEither: false, MinPurchase: 400, PaymentMethod: "cod"}},
		},
		{
			Name: "xdao25j200q", Code: "code145", IsEcommerce: true,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: DiscountVoucher{Amount: 25, Condition: Condition{IsEither: true, MinPurchase: 200, PaymentMethod: "qris"}},
		},
		{
			Name: "xdaa25j200q", Code: "code146", IsEcommerce: true,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: DiscountVoucher{Amount: 25, Condition: Condition{IsEither: false, MinPurchase: 200, PaymentMethod: "qris"}},
		},
		{
			Name: "xdao25j400q", Code: "code147", IsEcommerce: true,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: DiscountVoucher{Amount: 25, Condition: Condition{IsEither: true, MinPurchase: 400, PaymentMethod: "qris"}},
		},
		{
			Name: "xdaa25j400q", Code: "code148", IsEcommerce: false,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: DiscountVoucher{Amount: 25, Condition: Condition{IsEither: false, MinPurchase: 400, PaymentMethod: "qris"}},
		},
		{
			Name: "xdao25j200c", Code: "code149", IsEcommerce: true,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: DiscountVoucher{Amount: 25, Condition: Condition{IsEither: true, MinPurchase: 200, PaymentMethod: "cod"}},
		},
		{
			Name: "xdaa25j200c", Code: "code150", IsEcommerce: true,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: DiscountVoucher{Amount: 25, Condition: Condition{IsEither: false, MinPurchase: 200, PaymentMethod: "cod"}},
		},
		{
			Name: "xdao25j400c", Code: "code151", IsEcommerce: true,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: DiscountVoucher{Amount: 25, Condition: Condition{IsEither: true, MinPurchase: 400, PaymentMethod: "cod"}},
		},
		{
			Name: "xdaa25j400c", Code: "code152", IsEcommerce: true,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: DiscountVoucher{Amount: 25, Condition: Condition{IsEither: false, MinPurchase: 400, PaymentMethod: "cod"}},
		},
		{
			Name: "xdao25jt200q", Code: "code153", IsEcommerce: true,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: DiscountVoucher{Amount: 25, Condition: Condition{IsEither: true, MinPurchase: 200, PaymentMethod: "qris"}},
		},
		{
			Name: "xdaa25jt200q", Code: "code154", IsEcommerce: true,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: DiscountVoucher{Amount: 25, Condition: Condition{IsEither: false, MinPurchase: 200, PaymentMethod: "qris"}},
		},
		{
			Name: "xdao25jt400q", Code: "code155", IsEcommerce: true,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: DiscountVoucher{Amount: 25, Condition: Condition{IsEither: true, MinPurchase: 400, PaymentMethod: "qris"}},
		},
		{
			Name: "xdaa25jt400q", Code: "code156", IsEcommerce: false,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: DiscountVoucher{Amount: 25, Condition: Condition{IsEither: false, MinPurchase: 400, PaymentMethod: "qris"}},
		},
		{
			Name: "xdao25jt200c", Code: "code157", IsEcommerce: true,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: DiscountVoucher{Amount: 25, Condition: Condition{IsEither: true, MinPurchase: 200, PaymentMethod: "cod"}},
		},
		{
			Name: "xdaa25jt200c", Code: "code158", IsEcommerce: true,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: DiscountVoucher{Amount: 25, Condition: Condition{IsEither: false, MinPurchase: 200, PaymentMethod: "cod"}},
		},
		{
			Name: "xdao25jt400c", Code: "code159", IsEcommerce: true,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: DiscountVoucher{Amount: 25, Condition: Condition{IsEither: true, MinPurchase: 400, PaymentMethod: "cod"}},
		},
		{
			Name: "xdaa25jt400c", Code: "code160", IsEcommerce: true,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: DiscountVoucher{Amount: 25, Condition: Condition{IsEither: false, MinPurchase: 400, PaymentMethod: "cod"}},
		},
	}
}
