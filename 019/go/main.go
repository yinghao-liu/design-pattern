package main

import (
	"fmt"
	"observe/observe"
	"time"
)

type OB struct {
}

func (ob OB) Notify() {
	fmt.Printf("receive Notify\n")
}

// client客户端/使用者
func main() {
	var event observe.EventA
	event.Subscribe(new(OB))
	event.NotifyAll()
	time.Sleep(time.Second)
}
