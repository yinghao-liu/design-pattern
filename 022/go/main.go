package main

import "templatepattern/templatepattern"

func main() {
	var tt templatepattern.Template
	var con templatepattern.Concrete
	tt.Set(con)

	tt.Template()
}
