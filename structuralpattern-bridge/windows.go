package bridge

import "fmt"

type Windows struct {
	printer Printer
}

func (w *Windows) print() {
	fmt.Println("Print request for windows")
	w.printer.printFile()
}

func (w *Windows) setPrinter(printer Printer) {
	w.printer = printer
}
