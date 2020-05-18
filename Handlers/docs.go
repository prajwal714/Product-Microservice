// Package classification Petstore API.
//
// the purpose of this application is to provide an application
// that is using plain go code to define an API
//
// This should demonstrate all the possible comment annotations
// that are available to turn go code into a fully compliant swagger 2.0 spec
//
// Terms Of Service:
//
// there are no TOS at this moment, use at your own risk we take no responsibility
//
//     Schemes: http, https
//     Host: localhost:8080
//
//     BasePath: /
//     Version:	1.0.0
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

//	swagger:parameters updateProduct listSingleProduct deleteProducts
type productIDParameterWrapper struct {
	// The id of the products to update from the database
	//in: path
	//required: true
	ID int `json:"id"`
}

// No content
//	swagger:response noContent
type productsNoContent struct {
}
