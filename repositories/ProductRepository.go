package repositories

import (
	"github.com/Briskhorn/go-sample/crud-api/models"
	"github.com/google/uuid"
)

// ProductRepository is a structure that will get products from database
type ProductRepository struct {
}

// GetProductsRepo is a function of ProductRepository structure that will return all products from database
func (repository *ProductRepository) GetProductsRepo() (*[]models.Product, error) {
	// In this function, there should be a code that retreive data from a foreign source
	// here we return static list and it doesn't represent a real life case
	var result = []models.Product{
		{UUID: uuid.MustParse("98d0f0d6-ddf5-48fa-9e35-b09f40c294b9"), Name: "Ramen", Price: 10.00},
		{UUID: uuid.MustParse("c5614d55-7958-4696-934b-4e6acf105a1e"), Name: "Banh bao", Price: 4.00},
	}

	return &result, nil
}
