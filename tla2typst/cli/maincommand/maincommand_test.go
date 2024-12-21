package maincommand

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func CreateTestRootCommand() *cobra.Command {
	cmd := Command()
	return cmd
}

func TestMainCommand(t *testing.T) {
	testCases := []struct {
		testName   string
		args       []string
		shouldFail bool
		err        string
	}{
		{
			testName:   "Single filename",
			args:       []string{"fileName"},
			shouldFail: false,
		},
		{
			testName:   "Multiple filenames",
			args:       []string{"fileName1", "fileName2"},
			shouldFail: true,
		},
		{
			testName:   "No filenames",
			args:       []string{},
			shouldFail: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			cmd := CreateTestRootCommand()
			cmd.SetArgs(tc.args)

			err := cmd.Execute()

			switch tc.shouldFail {
			case false:
				assert.NoError(t, err, "Expected no error for a single filename.")
				assert.Equal(t, Filename, tc.args[0])
			case true:
				assert.Error(t, err, "Expected an error for no/multiple filenames")
			}
		})
	}
}
