package visitor

type shape interface {
	accept(Visitor)
}
