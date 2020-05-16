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

import data "microservice/Data"

// Generic error message returned as a string
// swagger:response errorResponse
type errorResponseWrapper struct {
	// Description of the error
	// in: body
	Body GenericError
}

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
