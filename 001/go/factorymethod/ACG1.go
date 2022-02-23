package factorymethod

import "fmt"

const ACG1 = "ACG1"

// concrete product
type ProductACG1 struct {
}

func (p ProductACG1) DoAC() {
	fmt.Printf("ProductACG1 DoAC\n")
}

// concrete factory
type FactoryACG1 struct {
}

func (f FactoryACG1) CreateProduct() ProductAC {
	return new(ProductACG1)
}

// package init
func init() {
	Factory.Register(ACG1, new(FactoryACG1))
}
