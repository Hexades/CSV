package csv

import (
	"log"
)

type Executable = func(data *Model, repo *Parser) Response

var BasicOpenFunc = func(data *Model, repo *Parser) *Response {
	log.Println("Received open: ", data)

	return NewResponse(data, nil)
}
var ReadFirstFunc = func(data *Model, repo *Parser) *Response {
	return NewResponse(data, nil)
}
