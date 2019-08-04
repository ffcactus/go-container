// Package service contains service components.package service
package service

import (
	"go-container/container"
	"go-container/message"
)

const (
	// BeanName is the name to get the service instance.
	BeanName = "Service"
)

var (
	messageService message.Message
)

func init() {
	container.Register(BeanName, func() interface{} {
		messageService = container.Bean(message.BeanName).(message.Message)
		return &DefaultImpl{}
	})
}

// Service define the interface.
type Service interface {
	Message() string
}

// DefaultImpl is the default implementation of Service.
type DefaultImpl struct {
}

// Message return the message.
func (s *DefaultImpl) Message() string {
	return messageService.Get()
}
