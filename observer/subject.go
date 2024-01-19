package observer

type Subject interface {
	Register(o Observer)
	DeRegister(o Observer)
	NotifyAll()
}
