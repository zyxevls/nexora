package model

type CartItem struct {
	ID        uint `gorm:"primaryKey"`
	CartID    uint
	ProductID uint
	Quantity  uint

	Product Product `gorm:"foreignKey:ProductID"`
}
