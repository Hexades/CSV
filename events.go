package csv

func NewEvent(value any) Event {
	return &EventModel{data: &Model{Value: value}}
}

type Event interface {
	Send(Response)
	Receive() Response
}
type EventModel struct {
	data            *Model
	executable      Executable
	responseChannel chan Response
}

func (e *EventModel) getChannel() chan Response {
	if e.responseChannel == nil {
		e.responseChannel = make(chan Response, 1)
	}
	return e.responseChannel
}

func (e *EventModel) Send(val Response) {
	e.getChannel() <- val
}

func (e *EventModel) Receive() Response {
	return <-e.getChannel()
}
func NewResponse(value any, err error) *Response {
	return &Response{&Model{value}, err}
}

type Response struct {
	data *Model
	Err  error
}
type Model struct {
	Value any
}
