package main

import "statepattern1/state"

func main() {
	var context state.OrderContext
	context.ChangeState(new(state.ToBePaid))
}
