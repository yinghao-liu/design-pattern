package abstractfactory

import (
	"fmt"
)

const G2 = "G2"

// concrete product
type ProductACG2 struct {
}

func (p ProductACG2) DoAC() {
	fmt.Printf("ProductACG2 DoAC\n")
}

// concrete product
type ProductTVG2 struct {
}

func (p ProductTVG2) DoTV() {
	fmt.Printf("ProductTVG2 DoTV\n")
}

// concrete factory
type FactoryG2 struct {
}

func (f FactoryG2) CreateProductAC() ProductAC {
	return new(ProductACG2)
}
func (f FactoryG2) CreateProductTV() ProductTV {
	return new(ProductTVG2)
}

// package init
func init() {
	Factory.Register(G2, new(FactoryG2))
}
