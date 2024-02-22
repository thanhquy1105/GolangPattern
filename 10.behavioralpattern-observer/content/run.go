package observer

func Run() {
	shirtItem := NewItem("Nike shirt")
	observerFirst := &Customer{ID: "avv@gmail.com"}
	observerSecond := &Customer{ID: "dsf@hotmail.com"}
	shirtItem.Register(observerFirst)
	shirtItem.Register(observerSecond)
	shirtItem.UpdateAvailability()
	shirtItem.DeRegister(observerSecond)
	shirtItem.UpdateAvailability()

}
