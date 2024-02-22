package visitor

import "fmt"

type AreaCalculator struct {
	area int
}

func (a *AreaCalculator) visitForCircle(c *Circle) {
	fmt.Println("calculating area for Circle")
}

func (a *AreaCalculator) visitForSquare(s *Square) {
	fmt.Println("calculating area for Square")
}

func (a *AreaCalculator) visitForRectangle(r *Rectangle) {
	fmt.Println("calculating area for Rectangle")
}
