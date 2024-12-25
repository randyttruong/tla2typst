package scanner

import (
	"fmt"

	"github.com/pkg/errors"
)

type ParserState struct {
	scanner *ScannerState
}

var (
	Parser = &ParserState{}
)

func InitParser(scanner *ScannerState) error {
	if scanner == nil {
		return errors.New("Scanner or stream does not exist")
	}

	Parser.scanner = scanner

	return nil
}

func (p *ParserState) ParseContent() error {

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
