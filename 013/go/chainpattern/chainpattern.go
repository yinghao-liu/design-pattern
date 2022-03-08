package chainpattern

import "fmt"

/************ HandleChain - context *******************/
type HandleChain struct {
	list []Handler
}

func (h HandleChain) Handle() {
	for i, j := range h.list {
		err := j.Handle()
		fmt.Printf("handle NO. %d ret:%v\n", i, err)
	}
}
func (h *HandleChain) AddHandler(hd Handler) {
	h.list = append(h.list, hd)
}

/************ abstract Handler *******************/
type Handler interface {
	Handle() error
}

/************ concrete  Handle *******************/
type ConcreteHandle1 struct {
}

func (c ConcreteHandle1) Handle() error {
	fmt.Printf("ConcreteHandle1\n")
	return nil
}

type ConcreteHandle2 struct {
}

func (c ConcreteHandle2) Handle() error {
	fmt.Printf("ConcreteHandle2\n")
	return nil
}
