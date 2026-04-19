package service

import (
	"errors"
	"nexora/internal/model"
	"nexora/internal/repository"

	"gorm.io/gorm"
)

var ErrProductNotFound = errors.New("product not found")
var ErrInvalidQuantity = errors.New("quantity must be greater than 0")

func GetOrCreateCart(userID uint) (*model.Cart, error) {
	cart, err := repository.GetCartByUserID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			cart = &model.Cart{UserID: userID}
			err = repository.CreateCart(cart)
			if err != nil {
				return nil, err
			}
			return cart, nil
		}
		return nil, err
	}
	return cart, nil
}

func AddToCart(userID uint, productID uint, qty int) error {
	if qty <= 0 {
		return ErrInvalidQuantity
	}

	_, err := repository.GetProductByID(productID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrProductNotFound
		}
		return err
	}

	cart, err := GetOrCreateCart(userID)
	if err != nil {
		return err
	}

	for _, item := range cart.Items {
		if item.ProductID == productID {
			item.Quantity += uint(qty)
			return repository.UpdateItem(&item)
		}
	}

	newItem := &model.CartItem{
		CartID:    cart.ID,
		ProductID: productID,
		Quantity:  uint(qty),
	}

	return repository.AddItem(newItem)
}

func GetCart(userID uint) (*model.Cart, error) {
	return GetOrCreateCart(userID)
}

func UpdateCartItem(itemID uint, qty int) error {
	item := &model.CartItem{
		ID:       itemID,
		Quantity: uint(qty),
	}
	return repository.UpdateItem(item)
}

func RemoveFromCart(itemID uint) error {
	return repository.DeleteItem(itemID)
}
