package main

import (
	"fmt"
	"singleton/singleton"
)

func main() {
	sin := singleton.GetInstance()
	fmt.Println(sin.ShowInstance())
	//a := new(singleton.singleton) //cannot refer to unexported name singleton.singleton
	//a.ShowInstance()
}
