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

func (r *Parser) OnEvent(parserChannel <-chan Event) {

	for evt := range parserChannel {
		log.Println("Execute event", evt)
		//evt.Execute(r)
	}
}
