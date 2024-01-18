package memento

import "fmt"

func Main() {
	careTaker := &CareTaker{
		mementoArray: make([]*Memento, 0),
	}
	originator := Originator{
		state: "A",
	}
	fmt.Printf("Originator current state: %s\n", originator.GetState())
	careTaker.AddMemento(originator.CreateMemento())

	originator.SetState("B")
	fmt.Printf("Originator current state: %s\n", originator.GetState())

	careTaker.AddMemento(originator.CreateMemento())
	originator.SetState("C")

	fmt.Printf("Originator current state: %s\n", originator.GetState())
	careTaker.AddMemento(originator.CreateMemento())

	originator.RestoreMemento(careTaker.GetMemento(1))
	fmt.Printf("Originator current state: %s\n", originator.GetState())

	originator.RestoreMemento(careTaker.GetMemento(0))
	fmt.Printf("Originator current state: %s\n", originator.GetState())
}
