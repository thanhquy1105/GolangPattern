package bridge

import "fmt"

type Mac struct {
	printer Printer
}

func (m *Mac) print() {
	fmt.Println("Print request for Mac")
	m.printer.printFile()
}

func (m *Mac) setPrinter(printer Printer) {
	m.printer = printer
}
