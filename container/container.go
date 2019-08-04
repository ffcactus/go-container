// Package container includes the files that represents the container concept.package container
package container

import (
	"fmt"
	"log"
)

var (
	order        []string
	beans        = make(map[string]interface{})
	constructors = make(map[string]func() interface{})
)

// Register will register the bean to the container and initialize it.
func Register(name string, constructor func() interface{}) {
	if beans[name] != nil {
		log.Printf("Registering duplicated bean: %s", name)
		return
	}
	order = append(order, name)
	constructors[name] = constructor
}

// Bean return bean's instance by bean's name.
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
