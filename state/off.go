package state

import "fmt"

type off struct {
}

func newOff() state {
	return &off{}
}

func (o *off) on(m *Machine) {
	fmt.Println("Machine is going from OFF to ON")
	m.setCurrent(newOn())
}

func (o *off) off(m *Machine) {
	fmt.Println("Machine already OFF")
}
