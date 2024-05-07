package csv

import (
	"log"
)

type Parser struct{}

func NewCSVParser() {
	p := new(Parser)
	log.Println("Created new parser: ", p)
	AddListener(p)
	log.Println("Added parser listener")
}

func (p *Parser) OnEvent(parserChannel <-chan Event) {

	for evt := range parserChannel {
		switch evt.(type) {
		case ReadEvent:
			{
				evt := evt.(ReadEvent)
				log.Println(evt)
				models := ReadAll(evt.FileName())
				evt.ResponseChannel() <- models
			}
			//case ReadCSVEvent:
			//	evt := evt.(ReadCSVEvent)
			//	log.Println(evt)
			//	evt.Channel() <- ReadAll(evt.Reader(), nil)
			//}

		}
	}
}
