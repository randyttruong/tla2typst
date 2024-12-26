package scanner

import (
	"fmt"

	"github.com/pkg/errors"
)

type Parser struct {
	scanner *ScannerState
	idx     int
}

var (
	Prsr = &Parser{idx: 0}
)

func InitParser(scanner *ScannerState) error {
	if scanner == nil {
		return errors.New("Scanner or stream does not exist")
	}

	Prsr.scanner = scanner

	return nil
}

func GetParser() *Parser {
	return Prsr
}

func (p *Parser) ParseContent() error {

	if p.scanner == nil {
		err := errors.New("Scanner or stream does not exist")
		return err
	}

	stream := p.scanner.stream

	for _, tok := range stream {
		fmt.Printf("%v", tok)
	}

	return nil
}
