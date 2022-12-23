package models

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID
	Name     string
	Surname  string
	Fullname string
}

func NewUser(name, surename string) *User {
	return &User{uuid.New(), name, surename, name + " " + surename}
}
