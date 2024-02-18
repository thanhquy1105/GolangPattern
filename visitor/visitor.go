package visitor

type Visitor interface {
	visitForSquare(s *Square)
	visitForCircle(c *Circle)
	visitForRectangle(r *Rectangle)
}
