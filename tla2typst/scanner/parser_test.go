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

// TODO: Should this
func Test_ASTGeneration(t *testing.T) {
	tests := []struct {
		testName   string
		testString string
		stream     []*Token
		expected   Expr
		err        error
	}{
		{
			testName:   "Simple variable declaration",
			testString: "Init == \\/ A /\\ B",
			stream: []*Token{
				{
					tokenType: IDENTIFIER,
					value:     "Init",
				},
				{
					tokenType: OPERATOR,
					value:     "==",
				},
				{
					tokenType: OPERATOR,
					value:     "\\/",
				},
				{
					tokenType: IDENTIFIER,
					value:     "A",
				},
				{
					tokenType: OPERATOR,
					value:     "/\\",
				},
				{
					tokenType: IDENTIFIER,
					value:     "B",
				},
				{
					tokenType: EOF,
					value:     "",
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.testName, func(t *testing.T) {
			s := GetScanner()
			p := GetParser()

			s.SetStream(tc.stream)
			InitParser(s)

			p.parse()
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
				op: &Token{
					tokenType: OPERATOR,
					value:     "==",
				},
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
