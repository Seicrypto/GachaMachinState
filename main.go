package main

import (
	"exmaple/sei/gachaMechineState/internal/gachaMachin"
	"exmaple/sei/gachaMechineState/internal/manager"
	"exmaple/sei/gachaMechineState/internal/orderState"
	"fmt"
	"sync"
)

func main() {
	gacha := gachaMachin.NewGachaContext()
	manager := manager.NewManager(gacha)

	var wg sync.WaitGroup
	for i := 0; i < 20; i++ { // 測試有20張訂單併發
		wg.Add(1)
		go func(orderNum int) {
			defer wg.Done()
			order := orderState.NewOrderContext()
			order.OrderTarget = "gacha"
			manager.HandleOrder(order)
			fmt.Printf("Order %d completed\n", orderNum)
		}(i)
	}
	wg.Wait()
}
