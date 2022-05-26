package state

import (
	"errors"
	"fmt"
	"time"
)

const (
	StateToBePaid     = "toBePaid"
	StateToBeReceived = "toBeReceived"
	StateDone         = "done"
	StateCanceling    = "canceling"
	StateCanceled     = "canceled"
)

const (
	OperationPay    = "pay"
	OperationCancel = "cancel"
)

var StateChan = make(chan bool) // 无实际用处，等完成退出

/*******************context************************/
type OrderContext struct {
	state OrderStater
}

func (o *OrderContext) ChangeState(state OrderStater) {
	o.state = state
	go o.state.Flow()
}
func (o *OrderContext) Handle(op string) error {
	ops := o.state.GetOperations()
	for _, j := range ops {
		if op == j {
			o.state.Handle(op)
			return nil
		}
	}
	return errors.New("not support operation")
}

/**********************state 接口***************************/
// abstract state
type OrderStater interface {
	Flow()                   // 状态流转
	Handle(op string)        // 接受操作
	GetOperations() []string // 获取当前状态可做的操作
}

/**********************state 具体实现***************************/
type OrderContextBase struct {
	order *OrderContext
}

func (b *OrderContextBase) SetContext(order *OrderContext) {
	b.order = order
}

// ToBePaid
type ToBePaid struct {
	OrderContextBase

	myTimer *time.Timer
}

func (s ToBePaid) Flow() {
	fmt.Printf("Flowing ToBePaid\n")
	s.myTimer = time.NewTimer(time.Second * 5)
	select {
	case a := <-s.myTimer.C:
		fmt.Printf("ToBePaid: time is up %+v\n", a)
		s.Handle(OperationCancel)
	}
	s.myTimer.Stop()
}
func (s ToBePaid) GetOperations() []string {
	return []string{OperationPay, OperationCancel}
}
func (s ToBePaid) Handle(op string) {
	switch op {
	case OperationPay:
		s.myTimer.Stop()

	case OperationCancel:
		s.myTimer.Stop()
		next := new(Canceling)
		next.SetContext(s.order)
		s.order.ChangeState(next)
	}
}

// ToBeReceived
type ToBeReceived struct {
	OrderContextBase
}

func (s ToBeReceived) Handle() {

}
func (s ToBeReceived) Flow() {
	fmt.Printf("Flowing ToBeReceived\n")
}
func (s ToBeReceived) GetOperations() []string {
	return []string{OperationPay, OperationCancel}
}

// Done
type Done struct {
	OrderContextBase
}

func (s Done) Flow() {
	fmt.Printf("Flowing Done\n")
	StateChan <- true
}
func (s Done) GetOperations() []string {
	return []string{OperationPay, OperationCancel}
}

func (s Done) Handle(op string) {

}

// Canceling
type Canceling struct {
	OrderContextBase
}

func (s Canceling) Flow() {
	fmt.Printf("Flowing Canceling\n")

	time.Sleep(time.Second * 2)

	next := new(Canceled)
	next.SetContext(s.order)
	s.order.ChangeState(next)

}
func (s Canceling) GetOperations() []string {
	return []string{OperationPay, OperationCancel}
}
func (s Canceling) Handle(op string) {

}

// Canceled
type Canceled struct {
	OrderContextBase
}

func (s Canceled) Flow() {
	fmt.Printf("Flowing Canceled\n")
	StateChan <- true
}
func (s Canceled) GetOperations() []string {
	return []string{OperationPay, OperationCancel}
}
func (s Canceled) Handle(op string) {

}
