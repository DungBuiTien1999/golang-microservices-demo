package data

import (
	"testing"
)

func TestChecksValidation(t *testing.T) {
	p := &Product{Name: "water", Price: 0.15, SKU: "abd-xys-hdg"}

	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
