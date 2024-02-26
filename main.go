package main

import (
	abstractfactory "github.com/thanhquy1105/GolangPattern/abstractfactory/content"
	adapter "github.com/thanhquy1105/GolangPattern/adapter/content"
	bridge "github.com/thanhquy1105/GolangPattern/bridge/content"
	builder "github.com/thanhquy1105/GolangPattern/builder/content"
	chainofresponsibility "github.com/thanhquy1105/GolangPattern/chainofresponsibility/content"
	command "github.com/thanhquy1105/GolangPattern/command/content"
	composite "github.com/thanhquy1105/GolangPattern/composite/content"
	iterator "github.com/thanhquy1105/GolangPattern/iterator/content"
	mediator "github.com/thanhquy1105/GolangPattern/mediator/content"
	memento "github.com/thanhquy1105/GolangPattern/memento/content"
	observer "github.com/thanhquy1105/GolangPattern/observer/content"
	prototype "github.com/thanhquy1105/GolangPattern/prototype/content"
	singleton "github.com/thanhquy1105/GolangPattern/singleton/content"
	state "github.com/thanhquy1105/GolangPattern/state/content"
	strategy "github.com/thanhquy1105/GolangPattern/strategy/content"
	templatemethod "github.com/thanhquy1105/GolangPattern/templatemethod/content"
	visitor "github.com/thanhquy1105/GolangPattern/visitor/content"
)

func main() {
	// 1. Singleton
	singleton.Run()

	//2. Builder
	builder.Run()

	//3. Abstract Factory
	abstractfactory.Run()

	//4. Prototype
	prototype.Run()

	// 5. Chain Of Responsibility
	chainofresponsibility.Run()

	//6. Command
	command.Run()

	//7. Iterator
	iterator.Run()

	//8. Mediator
	mediator.Run()

	//9. Memento
	memento.Run()

	//10. Observer
	observer.Run()

	//11. State
	state.Run()

	//12. Strategy
	strategy.Run()

	//13. Template method
	templatemethod.Run()

	//14. Visitor
	visitor.Run()

	//15. Adapter
	adapter.Run()

	//16. Bridge
	bridge.Run()

	//17. Composite
	composite.Run()
}
