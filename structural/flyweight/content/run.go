package flyweight

import (
	"fmt"
	"strconv"
)

func Run() {
	factory := NewFactory()

	for i := 1; i <= 15; i++ {
		starContext := NewContext(strconv.Itoa(i), 2)
		soldier := factory.CreateSoldier("foot man")
		soldier.Promote(starContext)
	}

	for i := 1; i <= 20; i++ {
		starContext := NewContext(strconv.Itoa(i), 3)
		soldier := factory.CreateSoldier("sea man")
		soldier.Promote(starContext)
	}

	fmt.Printf("Number of storage soldier map: %d\n", factory.GetSize())
}
