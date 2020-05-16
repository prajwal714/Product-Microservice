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

// ErrInvalidProductPath is an error message when the product path is not valid
var ErrInvalidProductPath = fmt.Errorf("Invalid Path, path should be /products/[id]")

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

// ValidationError is a collection of validation error messages
type ValidationError struct {
	Messages []string `json:"messages"`
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) MiddlewareProductValidation(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		p.l.Println("Middleware function called")
		prod := &data.Product{}

		err := data.FromJSON(prod, r.Body)

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
