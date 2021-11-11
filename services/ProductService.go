package services

import (
	"github.com/Briskhorn/go-sample/crud-api/interfaces"
	"github.com/Briskhorn/go-sample/crud-api/models"
)

// ProductService is a structure that will manage the data that was received by IProductRepository
type ProductService struct {
	interfaces.IProductRepository
}

// GetProductsService is a function of ProductService structure that will manage products received
// from IProductRepository when askink for all products
func (service *ProductService) GetProductsService() (*[]models.Product, error) {
	return service.GetProductsRepo()
}
