package main

import "decoratorpattern/decoratorpattern"

func main() {
	// before decorator
	var cal decoratorpattern.Calculate
	cal.Add()

	// use decorator
	var deco decoratorpattern.CalculateDecorate
	deco.SetCalculater(cal)
	deco.Add()
}
