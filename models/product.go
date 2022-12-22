package models

import (
	"errors"

	"github.com/google/uuid"
)

type Product struct {
	ID     uuid.UUID
	Amount uint
	Name   string
}

func NewProduct(name string) (*Product, error) {
	p := Product{uuid.New(), 1, ""}
	err := p.UpdateName(name)

	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (p *Product) UpdateName(name string) error {
	if len(name) < 1 {
		return errors.New("Product's name can't be empty")
	}

	p.Name = name
	return nil
}

func (p *Product) UpdateAmount(n uint) error {
	if n == 0 {
		return errors.New("Amount of product can't be less than 1")
	}
	p.Amount = n
	return nil
}
