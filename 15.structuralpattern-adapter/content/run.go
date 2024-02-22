package adapter

func Run() {
	client := &Client{}
	window := &Windows{}
	client.InsertSquareUSBInComputer(window)

	macAdapter := &MacAdapter{
		M: Mac{},
	}

	client.InsertSquareUSBInComputer(macAdapter)
}
