// Package controller includes the files for controller component.
package controller

import (
	"go-container/container"
	"go-container/service"
)

const (
	// BeanName can be used to get the instance of the controller.
	BeanName = "Controller"
)

var (
	svc service.Service
)

func init() {
	container.Register(BeanName, func() interface{} {
		svc = container.Bean(service.BeanName).(service.Service)
		return &DefaultImpl{}
	})
}

// Controller defines the interface.
type Controller interface {
	HTTPGet() string
}

// DefaultImpl implements the Controller interface.
type DefaultImpl struct {
}

// HTTPGet mock HTTP Get.
func (DefaultImpl) HTTPGet() string {
	return svc.Message()
}
