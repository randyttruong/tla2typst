package scanner

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
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

func Test_PrettyPrinting(t *testing.T) {
	tests := []struct {
		testName string
		ast      Expr
		expected string
		err      error
	}{
		{
			testName: "Init == 1",
			ast: &Binary{
				op: "==",
				left: &Identifier{
					value: "Init",
				},
				right: &Literal{
					value: "1",
				},
			},
			expected: "(== Init 1)",
		},
	}

	for _, tc := range tests {
		t.Run(tc.testName, func(t *testing.T) {

			var (
				p = &PrettyPrinter{}
			)

			res := fmt.Sprintf("%v", tc.ast.accept(p))

			assert.Equal(t, tc.expected, res, "Expected these to be the same")
		})
	}
}
