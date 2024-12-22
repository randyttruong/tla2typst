package scanner

import (
	"fmt"

	"github.com/pkg/errors"
)

// ScannerState represents the scanner's current state.
type ScannerState struct {
	loader *Loader
	Tok    []Token
	val    int
}

var (
	Scanner *ScannerState
)

// InitScanner passes the loader into ScannerState
func InitScanner(loader *Loader) error {
	if loader == nil || GetLoader().buf == nil {
		return errors.New("Loader or buffer does not exist/is empty.")
	}

	Scanner.loader = loader

	return nil
}

func GetScanner() *ScannerState {
	return Scanner
}

type ScannerStateDelta struct {
}
