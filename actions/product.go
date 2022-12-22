package actions

import (
	"bytes"
	"fmt"

	"github.com/kmisha/fan-in-pattern-go/models"
)

type Msg[T interface{}] struct {
	msg string
	o   T
}

// create
func CreateAction(ch chan<- string, name string) {
	var msg bytes.Buffer
	defer func() {
		ch <- msg.String()
	}()

	p, err := models.NewProduct(name)

	if err != nil {
		fmt.Fprintf(&msg, "error: can't create product with name = %s", name)
	} else {
		fmt.Fprintf(&msg, "create product with name = %s and id = %s", name, p.ID.String())
	}
}

// update
func UpdateAction(ch chan<- string, name string, p *models.Product) {
	var msg bytes.Buffer
	defer func() {
		ch <- msg.String()
	}()

	err := p.UpdateName(name)

	if err != nil {
		fmt.Fprintf(&msg, "error: can't update product's name %s", name)
	} else {
		fmt.Fprintf(&msg, "update product's name = %s and id = %s", name, p.ID.String())
	}
}
