package adapter

import "fmt"

type Windows struct {
}

func (w *Windows) insertInSquarePort() {
	fmt.Println("Insert square port into window machine")
}
