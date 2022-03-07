package main

import "chainpattern/chainpattern"

func main() {
	var chain chainpattern.HandleChain
	chain.AddHandler(new(chainpattern.ConcreteHandle1))
	chain.AddHandler(new(chainpattern.ConcreteHandle1))
	chain.Handle()
}
