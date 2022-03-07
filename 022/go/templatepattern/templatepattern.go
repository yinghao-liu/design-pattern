package templatepattern

import "fmt"

/* type Base struct {
}

func (b Base) Template() {
	b.Step1()
	b.Step2()
}

func (b Base) Step1() {
	fmt.Printf("base step 1\n")
}
func (b Base) Step2() {
	fmt.Printf("base step 2\n")
}

type Concrete struct {
	Base
}

func (b Concrete) Step2() {
	fmt.Printf("Concrete step 2\n")
} */

/*********************模板类*************************/
type Template struct {
	temp Templater
}

func (b *Template) Set(temp Templater) {
	b.temp = temp
}

func (b Template) Template() {
	b.temp.Step1()
	b.temp.Step2()
}

/*********************模板步骤接口*************************/
type Templater interface {
	Step1()
	Step2()
}

/*********************步骤基类*************************/
type Base struct {
}

func (b Base) Step1() {
	fmt.Printf("base step 1\n")
}
func (b Base) Step2() {
	fmt.Printf("base step 2\n")
}

/*********************步骤具体类*************************/
type Concrete struct {
	Base
}

func (b Concrete) Step2() {
	fmt.Printf("Concrete step 2\n")
}
