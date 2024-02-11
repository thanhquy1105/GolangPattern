package main

import (
	"github.com/thanhquy1105/GolangPattern/abstractfactory"
	"github.com/thanhquy1105/GolangPattern/builder"
	c "github.com/thanhquy1105/GolangPattern/chainofresponsibility"
	"github.com/thanhquy1105/GolangPattern/command"
	"github.com/thanhquy1105/GolangPattern/iterator"
	"github.com/thanhquy1105/GolangPattern/mediator"
	"github.com/thanhquy1105/GolangPattern/memento"
	"github.com/thanhquy1105/GolangPattern/observer"
	"github.com/thanhquy1105/GolangPattern/prototype"
	"github.com/thanhquy1105/GolangPattern/singleton"
	"github.com/thanhquy1105/GolangPattern/state"
	"github.com/thanhquy1105/GolangPattern/strategy"
	"github.com/thanhquy1105/GolangPattern/templatemethod"
)

func main() {
	// 1. Singleton
	singleton.Main()

	//2. Builder
	builder.Main()

	//3. Abstract Factory
	abstractfactory.Main()

	//4. Prototype
	prototype.Main()

	// 5. Chain Of Responsibility
	c.Main()

	//6. Command
	command.Main()

	//7. Iterator
	iterator.Main()

	//8. Mediator
	mediator.Main()

	//9. Memento
	memento.Main()

	//10. Observer
	observer.Main()

	//11. State
	state.Main()

	//12. Strategy
	strategy.Main()

	//13. Template method
	templatemethod.Main()
}
