package main

import (
	"fmt"
	"os"
)

type TV struct {
	Model string
}

func (t TV) On() {
	fmt.Fprintln(os.Stdout, t.Model, "Включен")
}

func (t TV) Off() {
	fmt.Fprintln(os.Stdout, t.Model, "выключен")
}

type DVDPleer struct {
	Model string
	DVD   string
}

func (d DVDPleer) On() {
	fmt.Fprintln(os.Stdout, d.Model, "Включен")
}

func (d DVDPleer) Off() {
	fmt.Fprintln(os.Stdout, d.Model, "выключен")
}

func (d DVDPleer) PlayMovie() {
	fmt.Fprintln(os.Stdout, "проигрываю диск ", d.DVD)
}

func (d DVDPleer) StopMovie() {
	fmt.Fprintln(os.Stdout, "останавливаю диск ", d.DVD)
}

type Subwoofer struct {
	Model string
}

func (s Subwoofer) On() {
	fmt.Fprintln(os.Stdout, s.Model, "Включен")
}

func (s Subwoofer) Off() {
	fmt.Fprintln(os.Stdout, s.Model, "выключен")
}

type HomeTheater struct {
	TV
	DVDPleer
	Subwoofer
}

func (h HomeTheater) On() {
	h.TV.On()
	h.DVDPleer.On()
	h.Subwoofer.On()
}

func (h HomeTheater) Off() {
	h.TV.Off()
	h.DVDPleer.Off()
	h.Subwoofer.Off()
}

func (h HomeTheater) PlayMovie() {
	h.DVDPleer.PlayMovie()
}

func (h HomeTheater) StopMovie() {
	h.DVDPleer.StopMovie()
}

func main() {
	H := HomeTheater{
		TV{
			Model: "Samsung",
		},
		DVDPleer{
			Model: "JVC",
			DVD:   "От заката до рассвета",
		},
		Subwoofer{
			Model: "JBL",
		},
	}
	H.On()
	H.PlayMovie()
	H.StopMovie()
	H.Off()
}
