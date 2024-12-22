package scanner

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testFile1 = "../tests/sample_tla_code/Bakery.tla"
	testFile2 = "../tests/sample_tla_code/Caesar.tla"
	testFile3 = "../tests/sample_tla_code/NoPermissions100.tla"
)

func TestLoadConfigNormal(t *testing.T) {
	tests := []struct {
		testName string
		filename string
		err      error
	}{
		{
			testName: "Bakery Algorithm",
			filename: testFile1,
		},
		{
			testName: "Caesar Consensus Algorithm",
			filename: testFile2,
		},
	}

	for _, tc := range tests {
		t.Run(tc.testName, func(t *testing.T) {

			err := LoadDocument(tc.filename)

			assert.NoError(t, err, "Expected no error for loading the file.\n")
		})
	}
}

func TestLoadConfigWithBadFilePermissions(t *testing.T) {
	tests := []struct {
		testName string
		filename string
		err      error
	}{
		{
			testName: "100 Permissions",
			filename: testFile3,
		},
	}

	for _, tc := range tests {
		t.Run(tc.testName, func(t *testing.T) {
			err := LoadDocument(tc.filename)

			assert.Error(t, err, "Expected an error for attempting to read a file with 100 file permissions.\n")
		})
	}
}
