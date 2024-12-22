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

func (s *ScannerState) GetBuffer() ([]string, error) {
	if Scanner.loader == nil || Scanner.loader.buf == nil {
		err := errors.New("Unable to get buffer, loader or bytes array does not exist. Exiting.")

		return nil, err
	}

	return Scanner.loader.buf, nil
}

func (s *ScannerState) ScanContent() error {

	buf, err := s.GetBuffer()

	if err != nil {
		return errors.Wrapf(err, "Something went wrong with scanning, got %v", err)
	}

	for _, tok := range buf {
		fmt.Printf("This is the current token: %v\n", tok)
	}

	return nil
}
