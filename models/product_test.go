package models

import (
	"testing"
)

func TestNewProduct(t *testing.T) {
	t.Run("we can create a product", func(t *testing.T) {
		p, err := NewProduct("Bike")

		if p == nil {
			t.Fatalf("we want a product with name = %s but got error = %q", "Bike", err)
		}
	})

	t.Run("we can't create a product with empty name", func(t *testing.T) {
		p, err := NewProduct("")

		if err == nil {
			t.Fatalf("we want want error = %q but got a product with name %s", err, p.Name)
		}
	})

	t.Run("amount of product should be 1", func(t *testing.T) {
		p, _ := NewProduct("Bike")

		if p.Amount != 1 {
			t.Fatalf("we want amount of product = 1 but got %d", p.Amount)
		}
	})
}

func TestUpdateName(t *testing.T) {
	t.Run("we can't set the empty name for product", func(t *testing.T) {
		p, _ := NewProduct("Bike")

		err := p.UpdateName("")

		if err == nil || p.Name == "" {
			t.Fatalf("we want an error, but got a product with empty name")
		}
	})
}

func TestUpdateAmount(t *testing.T) {
	t.Run("amount of product can't be less than one", func(t *testing.T) {
		p, _ := NewProduct("Bike")

		err := p.UpdateAmount(0)

		if err == nil || p.Amount < 1 {
			t.Fatalf("we want an error, but got a product with amount = %d", p.Amount)
		}
	})
}
