package main

import (
	"abstractfactory/abstractfactory"
	"fmt"
)

// abstract factory  client
func abstract(scheme string) {
	factory, err := abstractfactory.Factory.Get(scheme)
	if nil != err {
		fmt.Printf("err is %s\n", err.Error())
	}
	AC := factory.CreateProductAC()
	AC.DoAC()
	TV := factory.CreateProductTV()
	TV.DoTV()

}

// client
func main() {

	abstract("G1")

}
