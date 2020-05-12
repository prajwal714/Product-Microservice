package handlers

import (
	"log"
	data "microservice/Data"
	"net/http"
	"regexp"
	"strconv"
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
	//PUT method
	if r.Method == http.MethodPut {
		// expect the id in the URL
		rgx := regexp.MustCompile(`/([0-9]+)`)
		grp := rgx.FindAllStringSubmatch(r.URL.Path, -1)

		if len(grp) != 1 {
			http.Error(w, "Invalid URL", http.StatusBadRequest)
			return
		}

		if len(grp[0]) != 2 {
			p.l.Println("Invalid URL more than one grp capture")
			http.Error(w, "Invalid URL", http.StatusBadRequest)
			return
		}

		idString := grp[0][1]
		id, err := strconv.Atoi(idString)

		if err != nil {
			p.l.Println("Invalid URL unable to convert to number", idString)
			http.Error(w, "Invalid URL", http.StatusBadRequest)
			return
		}

		p.updateProducts(id, w, r)

	}

	//anything else Request
	// w.WriteHeader(http.StatusMethodNotAllowed)

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

func (p *Products) updateProducts(id int, w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT Request")

	prod := &data.Product{}

	err := prod.FromJSON(r.Body)

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
