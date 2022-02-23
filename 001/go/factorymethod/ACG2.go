package factorymethod

import "fmt"

const ACG2 = "ACG2"

// concrete product
type ProductACG2 struct {
}

func (p ProductACG2) DoAC() {
	fmt.Printf("ProductACG2 DoAC\n")
}

// concrete factory
type FactoryACG2 struct {
}

func (f FactoryACG2) CreateProduct() ProductAC {
	return new(ProductACG2)
}

// package init
func init() {
	Factory.Register(ACG2, new(FactoryACG2))
}
