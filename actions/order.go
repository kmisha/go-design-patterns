package actions

import (
	"bytes"
	"fmt"

	"github.com/kmisha/fan-in-pattern-go/models"
)

type CreateOrderAction struct {
	Owner *models.User
}

func (a *CreateOrderAction) Do() string {
	var msg bytes.Buffer
	o := models.NewOrder(a.Owner)

	fmt.Fprintf(&msg, "success: create an order with id = %s", o.ID.String())

	return msg.String()
}
