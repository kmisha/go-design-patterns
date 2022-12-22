package actions

import (
	"bytes"
	"fmt"

	"github.com/kmisha/fan-in-pattern-go/models"
)

const (
	CREATE_PRODUCT = "Create product"
	UPDATE_PRODUCT = "Update product"
	READ_PRODUCT   = "Read product"
	DELETE_PRODUCT = "Read product"
)

// Create a product with name
type CreateProductAction struct {
	name string
}

func (a *CreateProductAction) Do() string {
	var msg bytes.Buffer
	p, err := models.NewProduct(a.name)

	if err != nil {
		fmt.Fprintf(&msg, "error: can't create product with name = %s", a.name)
	} else {
		fmt.Fprintf(&msg, "success: create product with name = %s and id = %s", a.name, p.ID.String())
	}

	return msg.String()
}

// Update product name to
type UpdateProductAction struct {
	NewName string
	Product *models.Product
}

func (a *UpdateProductAction) Do() string {
	var msg bytes.Buffer

	err := a.Product.UpdateName(a.NewName)

	if err != nil {
		fmt.Fprintf(&msg, "error: can't update product's name %s", a.NewName)
	} else {
		fmt.Fprintf(&msg, "success: update product's name = %s and id = %s", a.NewName, a.Product.ID.String())
	}

	return msg.String()
}
