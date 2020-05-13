package handlers

import (
	"log"
	data "microservice/Data"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) GetProducts(w http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()

	err := lp.ToJSON(w)

	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}

}

func (p *Products) AddProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST products")

	//we create an empty product object
	prod := &data.Product{}
	//we decode the request body and feed it into the prod object
	err := prod.FromJSON(r.Body)

	if err != nil {
		http.Error(w, "Unable to unmarshal json", http.StatusBadRequest)
	}

	p.l.Printf("\n Prod:  %#v", prod)
	data.AddProduct(prod)
}
func (p *Products) UpdateProducts(w http.ResponseWriter, r *http.Request) { //PUT Request handler function

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Unable to convert id to integer", http.StatusBadRequest)
		return
	}
	p.l.Println("Handle PUT Request")

	prod := &data.Product{}

	err = prod.FromJSON(r.Body)

	if err != nil {
		http.Error(w, "Unable to unmarshal json", http.StatusBadRequest)
	}

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
