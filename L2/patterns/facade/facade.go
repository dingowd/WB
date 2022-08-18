package main

import (
	"fmt"
	"os"
)

type tv struct {
	model string
}

func (t tv) On() {
	fmt.Fprintln(os.Stdout, t.model, "Включен")
}

func (t tv) Off() {
	fmt.Fprintln(os.Stdout, t.model, "выключен")
}

func NewTV(m string) tv {
	return tv{
		model: m,
	}
}

type dvdPleer struct {
	model string
	dvd   string
}

func (d dvdPleer) On() {
	fmt.Fprintln(os.Stdout, d.model, "Включен")
}

func (d dvdPleer) Off() {
	fmt.Fprintln(os.Stdout, d.model, "выключен")
}

func (d dvdPleer) PlayMovie() {
	fmt.Fprintln(os.Stdout, "проигрываю диск ", d.dvd)
}

func (d dvdPleer) StopMovie() {
	fmt.Fprintln(os.Stdout, "останавливаю диск ", d.dvd)
}

func NewDVDPayer(m string) dvdPleer {
	return dvdPleer{
		model: m,
	}
}

type subwoofer struct {
	model string
}

func (s subwoofer) On() {
	fmt.Fprintln(os.Stdout, s.model, "Включен")
}

func (s subwoofer) Off() {
	fmt.Fprintln(os.Stdout, s.model, "выключен")
}

func Newsubwoofer(m string) subwoofer {
	return subwoofer{
		model: m,
	}
}

type HomeTheater struct {
	TV  tv
	DVD dvdPleer
	Sub subwoofer
}

func (h HomeTheater) On() {
	h.tv.On()
	h.dvdPleer.On()
	h.subwoofer.On()
}

func (h HomeTheater) Off() {
	h.tv.Off()
	h.dvdPleer.Off()
	h.subwoofer.Off()
}

func (h HomeTheater) PlayMovie() {
	h.dvdPleer.PlayMovie()
}

func (h HomeTheater) StopMovie() {
	h.dvdPleer.StopMovie()
}
