package memento

type CareTaker struct {
	mementoArray []*Memento
}

func (c *CareTaker) AddMemento(m *Memento) {
	c.mementoArray = append(c.mementoArray, m)
}

func (c *CareTaker) GetMemento(index int) *Memento {
	return c.mementoArray[index]
}
