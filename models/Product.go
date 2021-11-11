package models

import "github.com/google/uuid"

// Product is the structure that represent the product that can be sold
type Product struct {
	UUID  uuid.UUID `json:"uuid"`
	Name  string    `json:"name"`
	Price float32   `json:"price"`
}
