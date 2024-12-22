package scanner

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScanContent(t *testing.T) {

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
