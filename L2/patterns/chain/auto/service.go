package auto

type Service interface {
	Serve(*Car)
	SetNext(Service)
}

type Car struct {
	Name      string
	GetCar    bool
	ChangeOil bool
	WashCar   bool
}
