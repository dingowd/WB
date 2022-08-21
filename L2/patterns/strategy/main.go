package main

import (
	"bufio"
	"fmt"
	"os"
)

type Legion struct {
	Move
	Name string
}

func NewLegion() *Legion {
	return &Legion{}
}

type Move interface {
	Move()
}

type enemyIsWeek struct{}

func (e enemyIsWeek) Move() {
	fmt.Fprintln(os.Stdout, "Враг малочислен. Наступаем!!!")
}

func NewEnemyIsWeek() Move {
	return &enemyIsWeek{}
}

type enemyUnknown struct{}

func (e enemyUnknown) Move() {
	fmt.Fprintln(os.Stdout, "Враг неизвестен. Окапываемся!!!")
}

func NewEnemyUnknown() Move {
	return &enemyUnknown{}
}

type enemyIsStrong struct{}

func (e enemyIsStrong) Move() {
	fmt.Fprintln(os.Stdout, "Враг слишком силен. Отступаем!!!")
}

func NewEnemyIsStrong() Move {
	return &enemyIsStrong{}
}

func Scan() string {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	if err := in.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Input error:", err)
	}
	return in.Text()
}

func main() {
	legion := NewLegion()
	legion.Name = "Гордый римский легион"
	var intSay string
	fmt.Fprint(os.Stdout, "Что нам говорит разведка?:")
	intSay = Scan()
	switch intSay {
	case "Враг малочислен":
		legion.Move = NewEnemyIsWeek()
	case "Враг неизвестен":
		legion.Move = NewEnemyUnknown()
	case "Враг слишком силен":
		legion.Move = NewEnemyIsStrong()
	default:
		legion.Move = NewEnemyUnknown()
	}

	fmt.Fprintln(os.Stdout, legion.Name)
	legion.Move.Move()
}
