package main

import (
	"fmt"
	"os"
)

type HomeTheater interface {
	on()
	off()
	play()
	stop()
}

type Theater struct {
	isOn   bool
	isPlay bool
}

func (t *Theater) on() {
	t.isOn = true
	fmt.Fprintln(os.Stdout, "Home theater is on")
}

func (t *Theater) off() {
	if t.isOn {
		if t.isPlay {
			t.stop()
			t.isPlay = false
		}
		t.isOn = false
		fmt.Fprintln(os.Stdout, "Home theater is off")
	} else {
		fmt.Fprintln(os.Stdout, "You didn't turn on the home theater")
	}
}

func (t *Theater) play() {
	if t.isOn {
		t.isPlay = true
		fmt.Fprintln(os.Stdout, "let's enjoy the movie")
	} else {
		fmt.Fprintln(os.Stdout, "You didn't turn on the home theater")
	}
}

func (t *Theater) stop() {
	if t.isPlay {
		t.isPlay = false
		fmt.Fprintln(os.Stdout, "Stop the movie")
	} else {
		fmt.Fprintln(os.Stdout, "Nothing to stop")
	}
}
