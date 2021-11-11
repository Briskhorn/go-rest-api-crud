package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Briskhorn/go-sample/crud-api/controllers"
	"github.com/Briskhorn/go-sample/crud-api/repositories"
	"github.com/Briskhorn/go-sample/crud-api/services"
)

func handleRequests(controller *controllers.ProductController) {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8585"
	}
	http.HandleFunc("/product", controller.GetProducts)
	log.Default().Println("Running application on port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func main() {
	productRepository := &repositories.ProductRepository{}
	productService := &services.ProductService{IProductRepository: productRepository}
	productController := &controllers.ProductController{IProductService: productService}

	handleRequests(productController)
}
