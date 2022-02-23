package abstractfactory

import "errors"

type FactoryContext struct {
	context map[string]SuperFactory
}

func (f *FactoryContext) Init() {
	if nil == f.context {
		f.context = make(map[string]SuperFactory)
	}
}
func (f *FactoryContext) Register(scheme string, factory SuperFactory) {
	f.Init()
	f.context[scheme] = factory
}

func (f *FactoryContext) Get(scheme string) (SuperFactory, error) {
	if nil == f.context {
		return nil, errors.New("not support scheme")
	}

	if factory, exist := f.context[scheme]; exist {
		return factory, nil
	}
	return nil, errors.New("not support scheme")
}

var Factory FactoryContext

// abstract productAC
type ProductAC interface {
	DoAC()
}

// abstract productTV
type ProductTV interface {
	DoTV()
}

// abstract factory
type SuperFactory interface {
	CreateProductAC() ProductAC
	CreateProductTV() ProductTV
}
