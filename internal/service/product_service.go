package service

import "nexora/internal/repository"

func GetProducts() interface{} {
	products, err := repository.GetProducts()
	if err != nil {
		return nil
	}
	return products
}
