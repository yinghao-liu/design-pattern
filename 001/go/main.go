package main

import (
	"factorypattern/factorymethod"
	"factorypattern/simplefactory"
	"fmt"
)

// simple factory client
func simple(scheme string) {
	a := simplefactory.ChairFactory(scheme)
	a.Sit()
}

// factory method client
func method(scheme string) {
	factory, err := factorymethod.Factory.Get(scheme)
	if nil != err {
		fmt.Printf("err is %s\n", err.Error())
	}
	product := factory.CreateProduct()
	product.DoAC()

}

// client
func main() {
	// simple factory
	simple("blue")
	// factory method
	method(factorymethod.ACG1)

}
