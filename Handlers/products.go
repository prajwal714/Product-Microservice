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
