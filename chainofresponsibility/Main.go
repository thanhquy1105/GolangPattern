package chainofresponsibility

func Main() {
	cashier := &Cashier{}

	medical := &Medical{}
	medical.SetNext(cashier)

	doctor := &Doctor{}
	doctor.SetNext(medical)

	reception := &Reception{}
	reception.SetNext(doctor)

	patient := &Patient{Name: "You"}
	reception.Execute(patient)
}
