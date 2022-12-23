package actions

import (
	"bytes"
	"fmt"

	"github.com/kmisha/fan-in-pattern-go/models"
)

type CreateCartAction struct {
	owner   *models.User
	product *models.Product
}

func (a *CreateCartAction) Do() string {
	var msg bytes.Buffer
	o := models.NewCart(a.owner, a.product)

	fmt.Fprintf(&msg, "success: create an cart with id = %s", o.ID.String())

	return msg.String()
}
