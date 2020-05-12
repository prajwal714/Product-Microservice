package handlers

import (
	"log"
	data "microservice/Data"
	"net/http"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	//GET Method to getProducts and return all the products
	if r.Method == http.MethodGet {
		p.getProducts(w, r)
		return
	}

	//POST Method to post a new product
	if r.Method == http.MethodPost {
		p.addProduct(w, r)
		return
	}

	//anything else Request
	w.WriteHeader(http.StatusMethodNotAllowed)

}

func (p *Products) getProducts(w http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()

	err := lp.ToJSON(w)

	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}

}

func (p *Products) addProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST products")

	prod := &data.Product{}

	err := prod.FromJSON(r.Body)

	if err != nil {
		http.Error(w, "Unable to unmarshal json", http.StatusBadRequest)
	}

	p.l.Printf("\n Prod:  %#v", prod)
	data.AddProduct(prod)
}
