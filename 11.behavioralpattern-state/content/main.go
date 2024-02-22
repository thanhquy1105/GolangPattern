package state

func Run() {
	machine := NewMachine()
	machine.Off()
	machine.On()
	machine.On()
	machine.Off()
	machine.Off()
}
