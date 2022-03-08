package state

import (
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

/*******************context************************/
type OrderContext struct {
	state OrderStater
}

func (o *OrderContext) ChangeState(state OrderStater) {
	o.state = state
	o.state.Flow()
}
func (o *OrderContext) Pay(state OrderStater) {
	o.state = state
	o.state.Handle()
}

/**********************state***************************/
// abstract state
type OrderStater interface {
	Flow()
	Handle(op string)
	Operations() []string
}

// concrete state ToBePaid
type ToBePaid struct {
	order   OrderContext
	myTimer *time.Timer
}

func (s ToBePaid) Flow() {
	s.myTimer = time.NewTimer(time.Second * 5)
	select {
	case a := <-s.myTimer.C:
		fmt.Printf("time is up %+v\n", a)
	}

	s.myTimer.Stop()

}
func (s ToBePaid) Operations() []string {
	return []string{OperationPay, OperationCancel}
}
func (s ToBePaid) Handle(op string) {
	switch op {
	case OperationPay:
		s.myTimer.Stop()
	}
}

// concrete state ToBeReceived
type ToBeReceived struct {
}

func (s ToBeReceived) Handle() {

}

// concrete state Done
type Done struct {
}

func (s Done) Handle() {

}

// concrete state CancelRequest
type CancelRequest struct {
}

func (s CancelRequest) Handle() {

}

// concrete state Canceling
type Canceling struct {
}

func (s Canceling) Handle() {

}

// concrete state Canceled
type Canceled struct {
}

func (s Canceled) Handle() {

}
