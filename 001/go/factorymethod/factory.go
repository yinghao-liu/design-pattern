package factorymethod

import "errors"

type FactoryContext struct {
	context map[string]FactoryAC
}

func (f *FactoryContext) Init() {
	if nil == f.context {
		f.context = make(map[string]FactoryAC)
	}
}
func (f *FactoryContext) Register(scheme string, factory FactoryAC) {
	f.Init()
	f.context[scheme] = factory
}

func (f *FactoryContext) Get(scheme string) (FactoryAC, error) {
	if nil == f.context {
		return nil, errors.New("not support scheme")
	}

	if factory, exist := f.context[scheme]; exist {
		return factory, nil
	}
	return nil, errors.New("not support scheme")
}

var Factory FactoryContext

// abstract product
type ProductAC interface {
	DoAC()
}

// abstract factory
type FactoryAC interface {
	CreateProduct() ProductAC
}
