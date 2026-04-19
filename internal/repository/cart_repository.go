package repository

import (
	"nexora/config"
	"nexora/internal/model"
)

func GetCartByUserID(userID uint) (*model.Cart, error) {
	var cart model.Cart
	err := config.DB.Preload("Items.Product").Where("user_id = ?", userID).First(&cart).Error
	if err != nil {
		return nil, err
	}
	return &cart, nil
}

func CreateCart(cart *model.Cart) error {
	return config.DB.Create(cart).Error
}

func AddItem(cartItem *model.CartItem) error {
	return config.DB.Create(cartItem).Error
}

func UpdateItem(cartItem *model.CartItem) error {
	return config.DB.Save(cartItem).Error
}

func DeleteItem(cartItemID uint) error {
	return config.DB.Delete(&model.CartItem{}, cartItemID).Error
}

func ClearCart(cartID uint) error {
	return config.DB.Where("cart_id = ?", cartID).Delete(&model.CartItem{}).Error
}
