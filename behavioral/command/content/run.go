package command

func Run() {
	tv := &Tivi{}
	onCommand := &OnCommand{device: tv}

	offCommand := &OffCommand{device: tv}

	onButton := Button{
		command: onCommand,
	}
	onButton.press()
	offButton := Button{
		command: offCommand,
	}
	offButton.press()

}
