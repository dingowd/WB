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
