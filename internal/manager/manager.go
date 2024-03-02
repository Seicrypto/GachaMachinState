package manager

import (
	"exmaple/sei/gachaMechineState/internal/gachaMachin"
	"exmaple/sei/gachaMechineState/internal/orderState"
	"fmt"
)

type Manager struct {
	Gacha *gachaMachin.GachaContext
}

func NewManager(gacha *gachaMachin.GachaContext) *Manager {
	return &Manager{
		Gacha: gacha,
	}
}

func (oneself *Manager) HandleOrder(order *orderState.OrderContext) {
	switch order.OrderTarget {
	case "gacha":
		fmt.Println("Routing to Gacha machine")
		oneself.Gacha.Request()
	default:
		fmt.Println("Unknown order target")
	}
}
