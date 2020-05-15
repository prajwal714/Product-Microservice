package handlers

import (
	data "microservice/Data"
	"net/http"
)

//	swagger:route POST /products products addNewProduct
//	Adds a New Product in the list
//	responses:
//		200: productsResponse
func (p *Products) AddProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST products")

	//we create an empty product object
	prod := r.Context().Value(KeyProduct{}).(*data.Product)

	p.l.Printf("\n Prod:  %#v", prod)
	data.AddProduct(prod)
}
