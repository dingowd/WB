package main

import (
	"fmt"
	"os"
)

type DrinkingLoader struct {
	sober        State
	drunk        State
	hangover     State
	currentState State
	Name         string
}

func (l *DrinkingLoader) setState(s State) {
	l.currentState = s
}

func (l *DrinkingLoader) StartWork() error {
	return l.currentState.StartWork(l.Name)
}

func (l *DrinkingLoader) StopWork() error {
	return l.currentState.StopWork(l.Name)
}

func (l *DrinkingLoader) Sleep() error {
	return l.currentState.Sleep(l.Name)
}

func (l *DrinkingLoader) Drink() error {
	return l.currentState.Drink(l.Name)
}

func main() {
	v := &DrinkingLoader{
		Name:     "Джон",
		sober:    &Sober{},
		drunk:    &Drunk{},
		hangover: &Hangover{},
	}

	// sober
	v.setState(v.sober)
	fmt.Fprintln(os.Stdout, "SOBER*********************")
	if err := v.Sleep(); err != nil {
		fmt.Fprintln(os.Stdout, err.Error())
	}
	if err := v.StartWork(); err != nil {
		fmt.Fprintln(os.Stdout, err.Error())
	}
	if err := v.Drink(); err != nil {
		fmt.Fprintln(os.Stdout, err.Error())
	}
	if err := v.StopWork(); err != nil {
		fmt.Fprintln(os.Stdout, err.Error())
	}
	if err := v.Drink(); err != nil {
		fmt.Fprintln(os.Stdout, err.Error())
	}
	fmt.Fprintln(os.Stdout)

	// drunk
	v.setState(v.drunk)
	fmt.Fprintln(os.Stdout, "DRUNK*********************")
	if err := v.Sleep(); err != nil {
		fmt.Fprintln(os.Stdout, err.Error())
	}
	if err := v.StartWork(); err != nil {
		fmt.Fprintln(os.Stdout, err.Error())
	}
	if err := v.StopWork(); err != nil {
		fmt.Fprintln(os.Stdout, err.Error())
	}
	if err := v.Drink(); err != nil {
		fmt.Fprintln(os.Stdout, err.Error())
	}
	fmt.Fprintln(os.Stdout)

	// hangover
	fmt.Fprintln(os.Stdout, "HANGOVER*******************")
	v.setState(v.hangover)
	if err := v.Sleep(); err != nil {
		fmt.Fprintln(os.Stdout, err.Error())
	}
	if err := v.StartWork(); err != nil {
		fmt.Fprintln(os.Stdout, err.Error())
	}
	if err := v.Drink(); err != nil {
		fmt.Fprintln(os.Stdout, err.Error())
	}
	if err := v.StopWork(); err != nil {
		fmt.Fprintln(os.Stdout, err.Error())
	}
	if err := v.StartWork(); err != nil {
		fmt.Fprintln(os.Stdout, err.Error())
	}
	if err := v.StopWork(); err != nil {
		fmt.Fprintln(os.Stdout, err.Error())
	}
}
