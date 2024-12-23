package scanner

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func tokenEquals(a, b *Token) (bool, error) {

	if a.tokenType != b.tokenType && a.value != b.value {
		err := fmt.Errorf("Expected tokenType: %v,\nActual tokenType: %v,\nExpected value: %v,\nActual value: %v",
			a.tokenType,
			b.tokenType,
			a.value,
			b.value,
		)
		return false, err
	}

	return true, nil
}

func runTokenVerification(s *ScannerState, stream []*Token, t *testing.T) {
	fmt.Printf("This is the token stream:\n")
	for _, tok := range s.stream {
		fmt.Printf("%v,\n", *tok)
	}
	for idx, tok := range stream {
		res, err := tokenEquals(tok, s.stream[idx])
		assert.True(t, res, err)
	}
}

func TestScanContentWithFiles(t *testing.T) {

	// TODO: Use a smaller one, this is just for testing the output in terminal.
	tests := []struct {
		testName string
		filename string
		err      error
	}{
		{
			testName: "Bakery Algorithm (Short)",
			filename: smallTestFile1,
		},
		{
			testName: "Caesar Consensus (Short)",
			filename: smallTestFile2,
		},
	}

	for _, tc := range tests {
		t.Run(tc.testName, func(t *testing.T) {

			scanner := GetScanner()

			// 1. Ingest document
			err := LoadDocument(tc.filename)
			assert.NoError(t, err, "There should be no issues with loading the test algorithms.\n")

			// 2. Initialize the scanner
			err = InitScanner(GetLoader())
			assert.NoError(t, err, "There should be no issues with populating the loader buf.\n")

			// 3. Scan Content
			err = scanner.ScanContent()
			assert.NoError(t, err, "There should be no issues with scanning.\n")

			// TODO: Check the tokens
		})
	}
}

func TestScanContentSpecialChar(t *testing.T) {
	tests := []struct {
		testName string
		err      error
	}{
		{},
	}

	for _, tc := range tests {
		t.Run(tc.testName, func(t *testing.T) {

		})
	}
}

func TestScanContentComments(t *testing.T) {
	tests := []struct {
		testName       string
		test           string
		expectedStream []*Token
		err            error
	}{
		{
			testName: "Single-line block comments 1",
			test:     "(* HELLO *)",
			expectedStream: []*Token{
				&Token{
					tokenType: BLOCK_COMMENT,
					value:     "(* HELLO *)",
				},
			},
		},
		{
			testName: "Single-line block comments 2",
			test:     "(****************)",
			expectedStream: []*Token{
				&Token{
					tokenType: BLOCK_COMMENT,
					value:     "(****************)",
				},
			},
		},
		{
			testName: "Inline comments 1",
			test:     "Init == \\/ A /\\ B (* This is a comment *)",
			expectedStream: []*Token{
				&Token{
					tokenType: IDENTIFIER,
					value:     "INIT",
				},
				&Token{
					tokenType: OPERATOR,
					value:     "==",
				},
				&Token{
					tokenType: OPERATOR,
					value:     "\\/",
				},
				&Token{
					tokenType: IDENTIFIER,
					value:     "A",
				},
				&Token{
					tokenType: OPERATOR,
					value:     "/\\",
				},
				&Token{
					tokenType: IDENTIFIER,
					value:     "B",
				},
				&Token{
					tokenType: BLOCK_COMMENT,
					value:     "(* This is a comment *)",
				},
			},
		},
		{
			testName: "Inline comments 2",
			test:     "Init == \\/ A /\\ B \\* This is a comment",
			expectedStream: []*Token{
				&Token{
					tokenType: IDENTIFIER,
					value:     "INIT",
				},
				&Token{
					tokenType: OPERATOR,
					value:     "==",
				},
				&Token{
					tokenType: OPERATOR,
					value:     "\\/",
				},
				&Token{
					tokenType: IDENTIFIER,
					value:     "A",
				},
				&Token{
					tokenType: OPERATOR,
					value:     "/\\",
				},
				&Token{
					tokenType: IDENTIFIER,
					value:     "B",
				},
				&Token{
					tokenType: BLOCK_COMMENT,
					value:     "\\* This is a comment",
				},
			},
		},
		{
			testName: "Multi-line comments",
			test:     "(*\n\tThis is a test multiline comment.\n*)",
			expectedStream: []*Token{
				&Token{
					tokenType: BLOCK_COMMENT,
					value:     "(*\n\tThis is a test multiline comment.\n*)",
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.testName, func(t *testing.T) {

			scanner := GetScanner()
			scanner.stream = []*Token{}

			SetBuffer(tc.test)

			err := InitScanner(GetLoader())
			assert.NoError(t, err, "There should be no error with loading the buf.\n")

			err = scanner.ScanContent()
			assert.NoError(t, err, "There should be no issues with scanning.")

			fmt.Printf("This is the sample string:%v\n", tc.test)

			runTokenVerification(scanner, tc.expectedStream, t)

			// fmt.Printf("This is the token stream:\n")
			// for _, tok := range scanner.stream {
			// 	fmt.Printf("%v,\n", *tok)
			// }

		})
	}
}

func TestScanContentOnInlineCases(t *testing.T) {

	tests := []struct {
		testName       string
		sample         string
		expectedStream []*Token
		err            error
	}{
		{
			testName: "helloWorld",
			sample:   "print(\"Hello World!\")",
			expectedStream: []*Token{
				&Token{
					tokenType: IDENTIFIER,
					value:     "print",
				},
				&Token{
					tokenType: DELIMITER,
					value:     "(",
				},
				&Token{
					tokenType: STRING_LITERAL,
					value:     "\"Hello World!\"",
				},
				&Token{
					tokenType: DELIMITER,
					value:     ")",
				},
			},
		},
		{
			testName: "InitialState",
			sample:   "Init == \\/ A /\\ B\n        \\/ B /\\ C\n       \\/ C /\\ D\n",
			expectedStream: []*Token{
				&Token{
					tokenType: IDENTIFIER,
					value:     "Init",
				},
				&Token{
					tokenType: OPERATOR,
					value:     "==",
				},
				&Token{
					tokenType: OPERATOR,
					value:     "\\/",
				},
				&Token{
					tokenType: IDENTIFIER,
					value:     "A",
				},
				&Token{
					tokenType: OPERATOR,
					value:     "/\\",
				},
				&Token{
					tokenType: IDENTIFIER,
					value:     "B",
				},
				&Token{
					tokenType: OPERATOR,
					value:     "\\/",
				},
				&Token{
					tokenType: IDENTIFIER,
					value:     "B",
				},
				&Token{
					tokenType: OPERATOR,
					value:     "/\\",
				},
				&Token{
					tokenType: IDENTIFIER,
					value:     "C",
				},
				&Token{
					tokenType: OPERATOR,
					value:     "\\/",
				},
				&Token{
					tokenType: IDENTIFIER,
					value:     "C",
				},
				&Token{
					tokenType: OPERATOR,
					value:     "/\\",
				},
				&Token{
					tokenType: IDENTIFIER,
					value:     "D",
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.testName, func(t *testing.T) {

			scanner := GetScanner()
			scanner.stream = []*Token{}

			SetBuffer(tc.sample)

			err := InitScanner(GetLoader())
			assert.NoError(t, err, "There should be no error with loading the buf.\n")

			err = scanner.ScanContent()
			assert.NoError(t, err, "There should be no issues with scanning.")

			runTokenVerification(scanner, tc.expectedStream, t)
			fmt.Printf("This is the token stream:\n")
			for _, tok := range scanner.stream {
				fmt.Printf("%v,\n", *tok)
			}
		})
	}
}

func TestScanContentWithSpecificCases(t *testing.T) {

	// TODO: Use a smaller one, this is just for testing the output in terminal.
	tests := []struct {
		testName string
		filename string
		err      error
	}{
		{
			testName: "Set Comprehension",
			filename: setComprehension,
		},
		{
			testName: "Set Definition and Indexing",
			filename: setIndexing,
		},
		{
			testName: "Functional Operators",
			filename: functionalOperators,
		},
	}

	for _, tc := range tests {
		t.Run(tc.testName, func(t *testing.T) {

			scanner := GetScanner()

			// 1. Ingest document
			err := LoadDocument(tc.filename)
			assert.NoError(t, err, "There should be no issues with loading the test algorithms.\n")

			// 2. Initialize the scanner
			err = InitScanner(GetLoader())
			assert.NoError(t, err, "There should be no issues with populating the loader buf.\n")

			// 3. Scan Content
			err = scanner.ScanContent()
			assert.NoError(t, err, "There should be no issues with scanning.\n")

			// TODO: Check the tokens
		})
	}
}
