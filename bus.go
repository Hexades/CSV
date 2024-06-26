package csv

var bus = csvbus{}

type csvbus struct {
	listenerChannels []chan Event
}

func AddListener(eventListener EventListener) {
	listenerChannel := make(chan Event, 10)
	bus.listenerChannels = append(bus.listenerChannels, listenerChannel)
	go eventListener.OnEvent(listenerChannel)
}

func SendEvent(event Event) {
	for _, channel := range bus.listenerChannels {
		channel <- event
	}
}

type EventListener interface {
	OnEvent(eventChannel <-chan Event)
}
