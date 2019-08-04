# Overview


If you ever write code with Java Spring perhaps you do not quite understand the benefit it brings to you until you write code in something like Go. For example, you write controller in Java like this:
```java
@Controller
class Controller {

  @Autowired
  private MessageService message
  
  public String HttpGet() {
    return message.Get()
  }
}
```
Spring will help you create the Controller and set the message with an appropriate value, it will search the Classes that implements the MessageService interface, create it and assign it to message. message may have to be a singleton instance, in its constructor, a connection to the DB is made, so you should not create message again. Without Spring, you may need to create the message in the main() and pass it as an argument in Controller's constructor. However, it doesn't sound good, because:
- If the controller needs many others, all of them have to be passed in the constructor.
- The DB component and the controller component are not quite related, they both can work independently, why do you put DB component as an argument in the controller's constructor?
- The controller itself is a singleton, and the message service component may use other singleton object. So if we follow this style, we need to carefully create all the singletons at the beginning, it's hard to maintain for a large project.
Another benefit Spring brings is it make the code quite easy to test. You can replace the implementation in the UT without modifying the function.

How can we solve this problem with Go? We can mock what Spring does.

To do so, we need to create the singleton object in the right order. We can utilize the init() function.  Suppose you have controller.go and message.go and both of them have init() functions. Since controller.go will use Message interface, the init() function in message.go will be called first. If we insert an ID into a queue in the init() function and that ID represents the class, we would know the right order to initialize all the singletons.

Then we can initialize all the singletons in main(). Why not initialize the singletons directly in init()? Because that prevents us to replace the implementations in go test.

Here is the code, you can find the full code in the

https://github.com/ffcactus/go-container

Suppose we have controller, service and message components. The controller will call service, and service will call message.

The container:

```go
// container/container.go

// Package container includes the files that represents the container concept.package container
package container

import (
	"fmt"
)

var (
	// to record the order to constructor all the instances.
	order []string
	// saving all the instances.
	beans = make(map[string]interface{})
	// saving all the instances' constructor.
	constructors = make(map[string]func() interface{})
)

// Register will register the bean's name and constructor to the container
func Register(name string, constructor func() interface{}) {
	if beans[name] != nil {
		fmt.Printf("Registering duplicated bean: %s", name)
		return
	}
	order = append(order, name)
	constructors[name] = constructor
}

// Bean returns bean's instance by bean's name.
func Bean(name string) interface{} {
	return beans[name]
}

// Replace means instance, after this method the bean will be considered as initialized.
func Replace(name string, newInstance interface{}) {
	beans[name] = newInstance
}

// Init will initialize all the instances.
func Init() {
	for _, name := range order {
		fmt.Printf("Init bean %s\n", name)
		beans[name] = constructors[name]()
	}
}
```

The controller:

```go
// controller/controller.go

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
```

The service:

```go
// service/service.go

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
```

The message:

```go
// message/message.go

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
```

The controller's test:

```go
// controller/controller_test.go

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
```