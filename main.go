package main

import (
	"github.com/thanhquy1105/GolangPattern/abstractfactory"
	"github.com/thanhquy1105/GolangPattern/builder"
	"github.com/thanhquy1105/GolangPattern/singleton"
)

func main() {
	// 1. Singleton
	singleton.Main()

	//2. Builder
	builder.Main()

	//3. Abstract Factory
	abstractfactory.Main()
}
