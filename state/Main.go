package state

func Main() {
	machine := NewMachine()
	machine.Off()
	machine.On()
	machine.On()
	machine.Off()
	machine.Off()
}
