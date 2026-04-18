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
