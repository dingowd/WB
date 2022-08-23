package main

type Command interface {
	execute()
}

type OnCommand struct {
	theater HomeTheater
}

func (com *OnCommand) execute() {
	com.theater.on()
}

type OffCommand struct {
	theater HomeTheater
}

func (com *OffCommand) execute() {
	com.theater.off()
}

type PlayMovie struct {
	theater HomeTheater
}

func (com *PlayMovie) execute() {
	com.theater.play()
}

type StopMovie struct {
	theater HomeTheater
}

func (com *StopMovie) execute() {
	com.theater.stop()
}
