package state

import "fmt"

type on struct {
}

func newOn() state {
	return &on{}
}

func (o *on) on(m *Machine) {
	fmt.Println("Machine already on")
}

func (o *on) off(m *Machine) {
	fmt.Println("Machine is going from ON to OFF")
	m.setCurrent(newOff())
}
