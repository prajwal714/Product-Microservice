// Package Handlers for Products API.
//
// Documentation for Products API
//
// there are no TOS at this moment, use at your own risk we take no responsibility
//
//     Host: localhost
//     BasePath: /v
//     Version: 1.0.0
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package handlers

import (
	"context"
	"fmt"
	"log"
	data "microservice/Data"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Products struct {
	l *log.Logger
}
type KeyProduct struct{}

//	A list of products returs in the response
//	swagger:response productsResponse
type productsResponse struct {
	//	in: body
	Body []data.Product
}

//	swagger:parameters updateProduct
type productIDParameterWrapper struct {
	// The id of the products to update from the database
	//in: path
	//required: true
	ID int `json:"id"`
}

//	swagger:response noContent
type productsNoContect struct {
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

//	swagger:route GET /products products listProducts
//	Returns a list of products
//	responses:
//		200: productsResponse
func (p *Products) GetProducts(w http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()

	err := lp.ToJSON(w)

	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}

}

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

//	swagger:route PUT /products/{id} products updateProduct
//	Updates an existing product in the Products list
//	responses:
//		201: noContent

//	UpdateProcust updates an existing product in the database
func (p *Products) UpdateProducts(w http.ResponseWriter, r *http.Request) { //PUT Request handler function

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Unable to convert id to integer", http.StatusBadRequest)
		return
	}
	p.l.Println("Handle PUT Request")

	prod := r.Context().Value(KeyProduct{}).(*data.Product)

	err = data.UpdateProduct(id, prod)

	if err == data.ErrProductNotFound {
		http.Error(w, "Product Not Found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "Something went wrong", http.StatusBadRequest)
		return
	}

}

func (p *Products) MiddlewareProductValidation(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		p.l.Println("Middleware function called")
		prod := &data.Product{}

		err := prod.FromJSON(r.Body)

		if err != nil {
			p.l.Println("[ERROR] Error reading product", err)
			http.Error(w, "Error Reading Product", http.StatusBadRequest)
			return
		}

		err = prod.Validate()
		if err != nil {
			p.l.Println("[ERROR] Error Validating product", err)

			http.Error(w, fmt.Sprintf("Error Validating Product: %s", err), http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		req := r.WithContext(ctx)
		next.ServeHTTP(w, req)
	})
}
