// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new products API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for products API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	AddNewProduct(params *AddNewProductParams) (*AddNewProductOK, error)

	DeleteProducts(params *DeleteProductsParams) (*DeleteProductsCreated, error)

	ListProducts(params *ListProductsParams) (*ListProductsOK, error)

	ListSingleProduct(params *ListSingleProductParams) (*ListSingleProductOK, error)

	UpdateProduct(params *UpdateProductParams) (*UpdateProductCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  AddNewProduct Adds a New Product in the list
*/
func (a *Client) AddNewProduct(params *AddNewProductParams) (*AddNewProductOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddNewProductParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addNewProduct",
		Method:             "POST",
		PathPattern:        "/products",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddNewProductReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AddNewProductOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for addNewProduct: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteProducts Removes a product from the given list
*/
func (a *Client) DeleteProducts(params *DeleteProductsParams) (*DeleteProductsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteProductsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteProducts",
		Method:             "DELETE",
		PathPattern:        "/products/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteProductsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteProductsCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteProducts: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListProducts Returns a list of products
*/
func (a *Client) ListProducts(params *ListProductsParams) (*ListProductsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListProductsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listProducts",
		Method:             "GET",
		PathPattern:        "/products",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListProductsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListProductsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listProducts: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListSingleProduct Returns a single product
*/
func (a *Client) ListSingleProduct(params *ListSingleProductParams) (*ListSingleProductOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListSingleProductParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listSingleProduct",
		Method:             "GET",
		PathPattern:        "/products{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListSingleProductReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListSingleProductOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listSingleProduct: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateProduct Updates an existing product in the Products list
*/
func (a *Client) UpdateProduct(params *UpdateProductParams) (*UpdateProductCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateProductParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateProduct",
		Method:             "PUT",
		PathPattern:        "/products/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateProductReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateProductCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateProduct: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}