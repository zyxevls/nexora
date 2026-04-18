package repository

import (
	"nexora/config"
	"nexora/internal/model"
)

func GetProducts() ([]model.Product, error) {
	var products []model.Product
	err := config.DB.Find(&products).Error
	return products, err
}

func GetProductByID(productID uint) (*model.Product, error) {
	var product model.Product
	err := config.DB.First(&product, productID).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}
