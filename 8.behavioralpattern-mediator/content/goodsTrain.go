package mediator

import "fmt"

type GoodsTrain struct {
	mediator Mediator
}

func (g *GoodsTrain) RequestArrival() {
	if g.mediator.CanLand(g) {
		fmt.Println("Goods Train: Landing")
	} else {
		fmt.Println("Goods Train: Waiting")
	}
}

func (g *GoodsTrain) Departure() {
	fmt.Println("Goods Train: Leaving")
	g.mediator.NotifyFree()
}

func (g *GoodsTrain) PermitArrival() {
	fmt.Println("Goods Train: Arrival Permitted. Landing")
}
