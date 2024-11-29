package domain

type DiscountVoucher struct {
	VoucherID uint
	Rate      float32 `gorm:"precision:5;scale:2" json:"rate"`
	Amount    uint32  `gorm:"precision:4" json:"amount"`
	Condition `gorm:"embedded" json:"-"`
}
