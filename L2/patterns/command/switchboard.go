package main

type Switchboard struct {
	command Command
}

func (s *Switchboard) press() {
	s.command.execute()
}
