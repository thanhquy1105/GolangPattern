package builder

import (
	"fmt"
)

func Main() {
	normalBuilder := GetBuilder("normal")
	iglooBuilder := GetBuilder("igloo")

	director := NewDirector(normalBuilder)
	normalHouse := director.BuildHouse()

	fmt.Printf("Normal House Door Type: %s\n", normalHouse.GetDoorType())
	fmt.Printf("Normal House Window Type: %s\n", normalHouse.GetWindowType())
	fmt.Printf("Normal House Num Floor: %d\n", normalHouse.GetNumFloor())

	director.setBuilder(iglooBuilder)
	iglooHouse := director.BuildHouse()
	fmt.Printf("\nIgloo House Door Type: %s\n", iglooHouse.GetDoorType())
	fmt.Printf("Igloo House Window Type: %s\n", iglooHouse.GetWindowType())
	fmt.Printf("Igloo House Num Floor: %d\n", iglooHouse.GetNumFloor())

}
