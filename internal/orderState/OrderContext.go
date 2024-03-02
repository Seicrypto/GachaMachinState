package orderState

type OrderContext struct {
	State       OrderState
	OrderTarget string
}

func NewOrderContext() *OrderContext {
	return &OrderContext{
		State: &NewOrderState{},
	}
}

func (o *OrderContext) ChangeState(newState OrderState) {
	o.State = newState
}
