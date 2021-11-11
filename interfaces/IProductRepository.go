package interfaces

import "github.com/Briskhorn/go-sample/crud-api/models"

// IProductRepository is the interface that specify what it's imprementation should do
type IProductRepository interface {
	GetProductsRepo() (*[]models.Product, error)
}
