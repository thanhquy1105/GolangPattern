package visitor

type shape interface {
	Accept(Visitor)
}
