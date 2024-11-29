package domain

type FreeShippingVoucher struct {
	VoucherID uint
	Condition `gorm:"embedded" json:"-"`
}
