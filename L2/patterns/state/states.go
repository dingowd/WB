package main

import (
	"errors"
	"fmt"
	"os"
)

type State interface {
	Sleep(n string) error
	StartWork(n string) error
	StopWork(n string) error
	Drink(n string) error
}

type Sober struct {
	w, f bool
}

func (s *Sober) Sleep(n string) error {
	return errors.New(n + " не хочет спать, может работать")
}

func (s *Sober) StartWork(n string) error {
	s.w = true
	fmt.Fprintln(os.Stdout, n, "готов к работе")
	return nil
}

func (s *Sober) StopWork(n string) error {
	if s.w {
		s.f = true
		fmt.Fprintln(os.Stdout, n, "хорошо поработал сегодня")
		return nil
	}
	return errors.New(n + ", ты где был?")
}

func (s *Sober) Drink(n string) error {
	if s.f {
		fmt.Fprintln(os.Stdout, n, ", только чуть-чуть")
		return nil
	}
	return errors.New(n + ", работе еще не окончена")
}

type Drunk struct {
}

func (s *Drunk) Sleep(n string) error {
	fmt.Fprintln(os.Stdout, n, ", чтоб завтра был как огурчик, спи")
	return nil
}

func (s *Drunk) StartWork(n string) error {
	return errors.New(n + " спит")
}

func (s *Drunk) StopWork(n string) error {
	return errors.New(n + " спит")
}

func (s *Drunk) Drink(n string) error {
	return errors.New(n + ", только попробуй!!!")
}

type Hangover struct {
	w, h bool
}

func (s *Hangover) Sleep(n string) error {
	return errors.New(n + ", возьми себя в руки")
}

func (s *Hangover) StartWork(n string) error {
	if s.h {
		fmt.Fprintln(os.Stdout, n, ", руки не дрожат? Работай.")
		s.w = true
		return nil
	}
	return errors.New(n + " ты не можешь работать.")
}

func (s *Hangover) StopWork(n string) error {
	if s.w {
		fmt.Fprintln(os.Stdout, n, " хорошо поработал сегодня")
		return nil
	}
	return errors.New(n + ", ты где был?")
}

func (s *Hangover) Drink(n string) error {
	s.h = true
	fmt.Fprintln(os.Stdout, n, ", только чуть-чуть.")
	return nil
}
