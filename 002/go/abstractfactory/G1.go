package abstractfactory

import (
	"fmt"
)

const G1 = "G1"

// concrete product
type ProductACG1 struct {
}

func (p ProductACG1) DoAC() {
	fmt.Printf("ProductACG1 DoAC\n")
}

// concrete product
type ProductTVG1 struct {
}

func (p ProductTVG1) DoTV() {
	fmt.Printf("ProductTVG1 DoTV\n")
}

// concrete factory
type FactoryG1 struct {
}

func (f FactoryG1) CreateProductAC() ProductAC {
	return new(ProductACG1)
}
func (f FactoryG1) CreateProductTV() ProductTV {
	return new(ProductTVG1)
}

// package init
func init() {
	Factory.Register(G1, new(FactoryG1))
}
