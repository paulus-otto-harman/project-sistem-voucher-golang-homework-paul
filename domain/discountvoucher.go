package domain

type DiscountVoucher struct {
	ID        uint `gorm:"primaryKey;autoIncrement" json:"id"`
	VoucherID uint
	Rate      float32 `gorm:"precision:5;scale:2" json:"rate"`
	Amount    uint32  `gorm:"precision:4" json:"amount"`
	Condition `gorm:"embedded" json:"-"`
}
