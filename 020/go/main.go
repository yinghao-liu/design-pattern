package main

import (
	"fmt"
	"statepattern1/state"
)

func main() {
	var context state.OrderContext
	initState := new(state.ToBePaid)
	initState.SetContext(&context)

	context.ChangeState(initState)

	done := <-state.StateChan
	fmt.Printf("done:%v\n", done)
}
