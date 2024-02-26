package visitor

func Run() {
	square := &Square{Side: 2}
	circle := &Circle{Radius: 3}
	rectangle := &Rectangle{A: 4, B: 6}

	areaCal := &AreaCalculator{}
	square.Accept(areaCal)
	circle.Accept(areaCal)
	rectangle.Accept(areaCal)

	perimeterCal := &PerimeterCalculator{}
	square.Accept(perimeterCal)
	circle.Accept(perimeterCal)
	rectangle.Accept(perimeterCal)
}
