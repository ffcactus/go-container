// Package message includes the files for message component.
package message

import (
	"go-container/container"
)

const (
	// BeanName can be used to get the instance.
	BeanName = "Message"
)

func init() {
	container.Register(BeanName, func() interface{} {
		return &DefaultImpl{
			message: "Hello, World",
		}
	})
}

// Message is the interface for message component.
type Message interface {
	Get() string
}

// DefaultImpl implements the Message interface.
type DefaultImpl struct {
	message string
}

// Get returns the message.
func (impl *DefaultImpl) Get() string {
	return impl.message
}
