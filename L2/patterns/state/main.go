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

// Используется в тех случаях, когда во время выполнения программы объект должен
// менять своё поведение в зависимости от своего состояния

// Widget — класс, объекты которого должны менять своё поведение в зависимости от
// состояния.

// IState — интерфейс, который должен реализовать каждое из конкретных
// состояний. Через этот интерфейс объект Widget взаимодействует с состоянием,
// делегируя ему вызовы методов. Интерфейс должен содержать средства для обратной
// связи с объектом, поведение которого нужно изменить. Для этого используется
// событие (паттерн Publisher — Subscriber). Это необходимо для того, чтобы в
// процессе выполнения программы заменять объект состояния при появлении событий.
// Возможны случаи, когда сам Widget периодически опрашивает объект состояния на
// наличие перехода.

// StateA … StateZ — классы конкретных состояний. Должны
// содержать информацию о том, при каких условиях и в какие состояния может
// переходить объект из текущего состояния. Например, из StateA объект может
// переходить в состояние StateB и StateC, а из StateB — обратно в StateA и так
// далее. Объект одного из них должен содержать Widget при создании.

// Применение данного паттерна может быть затруднено, если состояния должны обмениваться
// данными, или одно состояние настраивает свойства другого. В этом случае
// понадобится глобальный объект, что не очень хорошее архитектурное решение.
