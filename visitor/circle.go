package visitor

type Circle struct {
	Radius int
}

func (s *Circle) Accept(v Visitor) {
	v.visitForCircle(s)
}
