package main

func main() {
	H := HomeTheater{
		TV:  NewTV("Samsung"),
		DVD: NewDVDPayer("JVC"),
		Sub: Newsubwoofer("JBL"),
	}

	H.On()
	H.PlayMovie()
	H.StopMovie()
	H.Off()
}

// Шаблон фасад (англ. Facade) — структурный шаблон проектирования, позволяющий
// скрыть сложность системы путём сведения всех возможных внешних вызовов к
// одному объекту, делегирующему их соответствующим объектам системы.
