package service

import (
	"context"
)

type EntityService struct {
	Updates chan string
}

type Action interface {
	Do() string
}

func NewEntityService(ctx context.Context) *EntityService {
	return &EntityService{make(chan string)}
}

func (s *EntityService) Do(act Action) {
	s.Updates <- act.Do()
}
