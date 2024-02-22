package command

import "fmt"

type Tivi struct {
	isRunning bool
}

func (t *Tivi) on() {
	t.isRunning = true
	fmt.Println("Turning Tivi on")
}

func (t *Tivi) off() {
	t.isRunning = true
	fmt.Println("Turning Tivi off")
}
