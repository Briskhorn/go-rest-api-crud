# [SAMPLE] [GO] [REST API] Crud-Api

This project is all about showing how for me a Go REST API should be organised
* controller -> manage http request and return a response with the correct code, body and header
* service -> implement the business aspect
* repository -> get raw data from a foreign source (other API or database)

## To run this project

### Prerequisites

To run this project you will first need to have 
* [Golang](https://golang.org/doc/install) installed on your computer

OR

* [Docker](https://www.docker.com/products/docker-desktop) installed on your computer

### Start the application

Using Go CLI:

* `go run main.go`

Using Docker:

* Build docker image: `docker build -t go-sample-project .`
* Run docker image: `docker run -e PORT=8585 -p 8585:8585 go-sample-project`

### Target Api

Get all products

* `http://localhost:8585/product`

