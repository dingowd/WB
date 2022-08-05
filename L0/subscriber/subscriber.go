package subscriber

type Subscriber interface {
	Start(stopChan chan struct{})
	Stop()
}
