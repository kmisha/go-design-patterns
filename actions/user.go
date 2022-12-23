package actions

import (
	"bytes"
	"fmt"

	"github.com/kmisha/fan-in-pattern-go/models"
)

type CreateUserAction struct {
	Name     string
	Surename string
}

func (a *CreateUserAction) Do() string {
	var msg bytes.Buffer
	u := models.NewUser(a.Name, a.Surename)

	fmt.Fprintf(&msg, "success: create a user with name = %s and id = %s", u.Name, u.ID.String())

	return msg.String()
}
