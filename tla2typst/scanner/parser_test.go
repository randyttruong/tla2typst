package scanner

import (
	"fmt"
	"testing"
)

func Test_parser_init(t *testing.T) {
	tests := []struct {
		testName string
		err      error
	}{}

	for _, tc := range tests {
		t.Run(tc.testName, func(t *testing.T) {
			fmt.Println("Hello world!")
		})
	}
}
