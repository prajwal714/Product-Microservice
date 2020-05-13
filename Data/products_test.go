package data

import "testing"

func TestChecksValidation(t *testing.T) {

	p := &Product{
		//Name:  "abc",
		Price: 4,
		SKU:   "abs-abs-abs",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}

}
