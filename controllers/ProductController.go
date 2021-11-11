package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Briskhorn/go-sample/crud-api/interfaces"
)

// ProductController is a structure that will manage http connection for product API
type ProductController struct {
	interfaces.IProductService
}

// GetProducts is a function of ProductController structure that will manage all requests to get all products
func (productController *ProductController) GetProducts(res http.ResponseWriter, req *http.Request) {

	products, err := productController.GetProductsService()
	if err != nil {
		//Handle error
	}

	json.NewEncoder(res).Encode(products)
}
