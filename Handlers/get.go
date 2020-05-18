package handlers

import (
	data "microservice/Data"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//	swagger:route GET /products products listProducts
//	Returns a list of products
//	responses:
//		200: productsResponse
func (p *Products) GetProducts(w http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()

	err := data.ToJSON(lp, w)

	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}

}

//	swagger:route GET /products{id} products listSingleProduct
//	Returns a single product
//	responses:
//		200: productsResponse
//		404: errorResponse

//Listsingle handles GET requests
func (p *Products) GetProductsByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, "Ubale to convert string to integer", http.StatusInternalServerError)
		return
	}

	p.l.Println("Get Product by ID request")

	prod, err := data.GetProductsByID(id)

	if err == data.ErrProductNotFound {
		p.l.Println("ERR Product not found")
		http.Error(w, "Error Product not found", http.StatusBadRequest)
		return
	}

	if err != nil {
		p.l.Println("Something went wrong")
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	err = data.ToJSON(prod, w)

}
