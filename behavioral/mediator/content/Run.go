package mediator

func Run() {
	stationManager := NewStationManager()
	passengerTrain := &PassengerTrain{
		mediator: stationManager,
	}

	goodsTrain := &GoodsTrain{
		mediator: stationManager,
	}

	passengerTrain.RequestArrival()
	goodsTrain.RequestArrival()
	passengerTrain.Departure()
}
