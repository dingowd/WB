package main

func main() {
	theater := &Theater{}

	onCommand := &OnCommand{
		theater: theater,
	}
	offCommand := &OffCommand{
		theater: theater,
	}
	playCommand := &PlayMovie{
		theater: theater,
	}
	stopCommand := &StopMovie{
		theater: theater,
	}
	on := &Switchboard{
		command: onCommand,
	}
	off := &Switchboard{
		command: offCommand,
	}
	play := &Switchboard{
		command: playCommand,
	}
	stop := &Switchboard{
		command: stopCommand,
	}

	on.press()
	off.press()
	play.press()
	on.press()
	play.press()
	off.press()
	stop.press()
}
