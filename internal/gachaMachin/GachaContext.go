package gachaMachin

import (
	"fmt"
	"sync"
)

type State interface {
	HandleRequest(*GachaContext)
}

type GachaContext struct {
	currentState State
	itemCount    int

	// States
	idleState        State
	dispensingState  State
	soldOutState     State
	maintenanceState State
	mu               sync.Mutex
}

// Initial new gacha state.
func NewGachaContext() *GachaContext {
	// Start with 10 items.
	context := &GachaContext{itemCount: 10}
	idle := &IdleState{}
	dispensing := &DispensingState{}
	soldOut := &SoldOutState{}
	maintenance := &MaintenanceState{}

	context.idleState = idle
	context.dispensingState = dispensing
	context.soldOutState = soldOut
	context.maintenanceState = maintenance
	context.currentState = idle

	return context
}

func (oneself *GachaContext) ChangeState(newState State) {
	oneself.currentState = newState
}

func (oneself *GachaContext) Request() {
	oneself.mu.Lock()
	defer oneself.mu.Unlock()
	oneself.currentState.HandleRequest(oneself)
}

func (oneself *GachaContext) addItem(count int) {
	oneself.itemCount += count
	fmt.Printf("Added %d items. Total items: %d.\n", count, oneself.itemCount)
	if _, ok := oneself.currentState.(*MaintenanceState); ok {
		oneself.ChangeState(oneself.idleState)
	}
}

func (oneself *GachaContext) dispenseItem() error {
	if oneself.itemCount > 0 {
		oneself.itemCount--
		fmt.Printf("Dispensed an item. Items left: %d.\n", oneself.itemCount)
		return nil
	} else {
		oneself.ChangeState(oneself.soldOutState)
		return fmt.Errorf("no items left to dispense")
	}
}
