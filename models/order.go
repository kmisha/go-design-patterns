package models

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID      uuid.UUID
	Owner   *User
	Created time.Time
	Ready   bool
}

func NewOrder(owner *User) *Order {
	return &Order{
		ID:      uuid.New(),
		Owner:   owner,
		Created: time.Now(),
		Ready:   false,
	}
}
