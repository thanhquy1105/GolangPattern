package bridge

// Bridge actually is a combination of Template Method and Strategy
func Run() {
	hpPrinter := &HP{}
	epsonPrinter := &Epson{}

	macComputer := &Mac{}
	macComputer.setPrinter(hpPrinter)
	macComputer.print()

	macComputer.setPrinter(epsonPrinter)
	macComputer.print()

	windowsComputer := &Windows{}
	windowsComputer.setPrinter(hpPrinter)
	windowsComputer.print()

	windowsComputer.setPrinter(epsonPrinter)
	windowsComputer.print()

}
