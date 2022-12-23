package models

import (
	"time"

	"github.com/google/uuid"
)

type Cart struct {
	ID      uuid.UUID
	Owner   *User
	Created time.Time
	Product *Product
}

func NewCart(owner *User, product *Product) *Cart {
	return &Cart{
		ID:      uuid.New(),
		Owner:   owner,
		Created: time.Now(),
		Product: product,
	}
}
