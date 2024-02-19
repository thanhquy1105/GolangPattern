package adapter

func Main() {
	client := &Client{}
	window := &Windows{}
	client.InsertSquareUSBInComputer(window)

	macAdapter := &MacAdapter{
		M: Mac{},
	}

	client.InsertSquareUSBInComputer(macAdapter)
}
