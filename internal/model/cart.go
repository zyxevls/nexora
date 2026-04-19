package model

type Cart struct {
	ID     uint `gorm:"primaryKey"`
	UserID uint
	Items  []CartItem
}
