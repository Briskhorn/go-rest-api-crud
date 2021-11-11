package interfaces

import "github.com/Briskhorn/go-sample/crud-api/models"

// IProductService is the interface that specify what it's imprementation should do
type IProductService interface {
	GetProductsService() (*[]models.Product, error)
}
