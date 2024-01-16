package main

import (
	"github.com/thanhquy1105/GolangPattern/abstractfactory"
	"github.com/thanhquy1105/GolangPattern/builder"
	c "github.com/thanhquy1105/GolangPattern/chainofresponsibility"
	"github.com/thanhquy1105/GolangPattern/command"
	"github.com/thanhquy1105/GolangPattern/iterator"
	"github.com/thanhquy1105/GolangPattern/mediator"
	"github.com/thanhquy1105/GolangPattern/prototype"
	"github.com/thanhquy1105/GolangPattern/singleton"
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
}
