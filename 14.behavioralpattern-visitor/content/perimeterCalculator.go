package visitor

import "fmt"

type PerimeterCalculator struct {
	perimeter int
}

func (p *PerimeterCalculator) visitForCircle(c *Circle) {
	fmt.Println("calculating perimeter for Circle")
}

func (p *PerimeterCalculator) visitForSquare(s *Square) {
	fmt.Println("calculating perimeter for Square")
}

func (p *PerimeterCalculator) visitForRectangle(r *Rectangle) {
	fmt.Println("calculating perimeter for Rectangle")
}
