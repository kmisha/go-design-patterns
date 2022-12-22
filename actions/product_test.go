package actions

import (
	"strings"
	"testing"

	"github.com/kmisha/fan-in-pattern-go/models"
)

// create
func TestCreateProduct(t *testing.T) {
	t.Run("we got a message when we creating a new product", func(t *testing.T) {
		action := CreateProductAction{name: "Bike"}

		got := action.Do()

		if !strings.HasPrefix(got, "success") {
			t.Fatalf("want positive result, but got %s", got)
		}

	})

	t.Run("we got an error message when we creating a new product with empty name", func(t *testing.T) {
		action := CreateProductAction{name: ""}

		got := action.Do()

		if !strings.HasPrefix(got, "error") {
			t.Fatalf("didn't get a message when create a product %v", got)
		}
	})
}

// update
func TestUpdateProduct(t *testing.T) {
	t.Run("we got a message when we creating a new product", func(t *testing.T) {
		p, _ := models.NewProduct("Bike")
		action := UpdateProductAction{NewName: "Super Bike", Product: p}

		got := action.Do()

		if !strings.HasPrefix(got, "success") {
			t.Fatalf("want positive result, but got %s", got)
		}
	})

	t.Run("we got an error message when we creating a new product with empty name", func(t *testing.T) {
		p, _ := models.NewProduct("Bike")
		action := UpdateProductAction{NewName: "", Product: p}

		got := action.Do()

		if !strings.HasPrefix(got, "error") {
			t.Fatalf("didn't get a message when create a product %v", got)
		}
	})
}
