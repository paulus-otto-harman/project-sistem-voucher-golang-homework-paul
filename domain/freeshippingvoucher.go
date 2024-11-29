package domain

type FreeShippingVoucher struct {
	ID        uint `gorm:"primaryKey;autoIncrement" json:"id"`
	VoucherID uint
	Condition `gorm:"embedded" json:"-"`
}
