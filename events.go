package csv

type Event interface {
}
type ReadEvent interface {
	Event
	FileName() string
	DataModel() any
	ResponseChannel() chan any
}
type readModel struct {
	fileName        string
	dataModel       any
	responseChannel chan any
}

func NewReadFileEvent(fileName string, dataModel any) ReadEvent {
	return &readModel{fileName, dataModel, nil}

}

func (e *readModel) ResponseChannel() chan any {
	if e.responseChannel == nil {
		e.responseChannel = make(chan any, 1)
	}
	return e.responseChannel
}

func (e *readModel) FileName() string {
	return e.fileName
}

func (e *readModel) DataModel() any {
	return e.dataModel
}
