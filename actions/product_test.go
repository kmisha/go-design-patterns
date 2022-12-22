package actions

import (
	"strings"
	"testing"
	"time"

	"github.com/kmisha/fan-in-pattern-go/models"
)

// create
func TestCreateProduct(t *testing.T) {
	t.Run("we got a message when we creating a new product", func(t *testing.T) {
		productCh := make(chan string)

		go CreateAction(productCh, "Bike")
		var got string
		select {
		case m := <-productCh:
			got = m
		case <-time.After(10 * time.Millisecond):
			got = ""
		}

		if got == "" {
			t.Fatalf("didn't get a message when create a product %v", got)
		}
	})

	t.Run("we got an error message when we creating a new product with empty name", func(t *testing.T) {
		productCh := make(chan string)

		go CreateAction(productCh, "")
		var got string
		select {
		case m := <-productCh:
			got = m
		case <-time.After(10 * time.Millisecond):
			got = ""
		}

		if !strings.HasPrefix(got, "error") {
			t.Fatalf("didn't get a message when create a product %v", got)
		}
	})
}

// update
func TestUpdateProduct(t *testing.T) {
	t.Run("we got a message when we creating a new product", func(t *testing.T) {
		productCh := make(chan string)

		p, _ := models.NewProduct("Bike")

		go UpdateAction(productCh, "Super bike", p)
		var got string
		select {
		case m := <-productCh:
			got = m
		case <-time.After(10 * time.Millisecond):
			got = ""
		}

		if got == "" {
			t.Fatalf("didn't get a message when create a product %v", got)
		}
	})

	t.Run("we got an error message when we creating a new product with empty name", func(t *testing.T) {
		productCh := make(chan string)

		p, _ := models.NewProduct("Bike")

		go UpdateAction(productCh, "", p)
		var got string
		select {
		case m := <-productCh:
			got = m
		case <-time.After(10 * time.Millisecond):
			got = ""
		}

		if !strings.HasPrefix(got, "error") {
			t.Fatalf("didn't get a message when create a product %v", got)
		}
	})
}
