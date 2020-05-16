package handlers

import (
	data "microservice/Data"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

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
