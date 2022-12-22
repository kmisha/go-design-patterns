package product

import (
	"context"
	"testing"
	"time"
)

type MockAction struct{}

const MOCK_MESSAGE = "Test Message"

func (a *MockAction) Do() string {
	time.Sleep(20 * time.Millisecond)
	return MOCK_MESSAGE
}

func TestProductService(t *testing.T) {
	service := NewEntityService(context.TODO())
	action := MockAction{}

	go func() {
		service.Do(&action)
	}()

	got := <-service.updates

	if got != MOCK_MESSAGE {
		t.Fatalf("want message = %s; got %s", MOCK_MESSAGE, got)
	}

}
