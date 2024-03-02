package gachaMachin

import "fmt"

type IdleState struct{}
type DispensingState struct{}
type SoldOutState struct{}
type MaintenanceState struct{}

func (oneself *IdleState) HandleRequest(context *GachaContext) {
	fmt.Println("Gacha is idle. Ready to dispense.")
	context.ChangeState(context.dispensingState)
}

func (oneself *DispensingState) HandleRequest(context *GachaContext) {
	err := context.dispenseItem()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		context.ChangeState(context.idleState)
	}
}

func (oneself *SoldOutState) HandleRequest(context *GachaContext) {
	fmt.Println("Gacha is sold out. Enter maintenance to add items.")
	context.ChangeState(context.maintenanceState)
}

func (oneself *MaintenanceState) HandleRequest(context *GachaContext) {
	fmt.Println("Gacha is under maintenance. Adding items...")
	context.addItem(10)
}
