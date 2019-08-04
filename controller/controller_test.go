package controller

import (
	"testing"
)

const (
	expected = "mock Message"
)

type mockService struct {
}

// Message return the message.
func (s *mockService) Message() string {
	return expected
}

func TestController_HTTPGet(t *testing.T) {
	svc = &mockService{}
	c := &DefaultImpl{}
	actual := c.HTTPGet()
	if expected != actual {
		t.Errorf("expect %v, actual %v", expected, actual)
	}
}
