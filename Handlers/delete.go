package handlers

import (
	data "microservice/Data"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//	swagger:route DELETE /products/{id} products deleteProducts
//	Removes a product from the given list
//	responses:
//		201: noContent

func (p *Products) DeleteProducts(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Unable to Convert id to integer", http.StatusInternalServerError)
		return
	}

	p.l.Println("Handle DELETE Request")

	err = data.DeleteProduct(id)

	if err == data.ErrProductNotFound {
		http.Error(w, "Product Not Found", http.StatusBadRequest)
		return

	}

	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

}
