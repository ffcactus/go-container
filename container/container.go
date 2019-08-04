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
