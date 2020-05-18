package main

import (
	"fmt"
	"microservice/client/client"
	"microservice/client/client/products"
	"testing"
)

func TestOurClient(t *testing.T) {
	cfg := client.DefaultTransportConfig().WithHost("localhost:8080")
	c := client.NewHTTPClientWithConfig(nil, cfg)
	params := products.NewListProductsParams()
	prods, err := c.Products.ListProducts(params)

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(prods)

}
