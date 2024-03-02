package orderState

import "fmt"

type OrderState interface {
	HandleRequest(*OrderContext)
}

type NewOrderState struct{}
type PaidState struct{}
type CancelledState struct{}

func (oneself *NewOrderState) HandleRequest(ctx *OrderContext) {
	fmt.Println("Initailed a order.")
}

func (oneself *PaidState) HandleRequest(ctx *OrderContext) {
	fmt.Println("Paid.")
}

func (oneself *CancelledState) HandleRequest(ctx *OrderContext) {
	fmt.Println("Canceled.")
}
