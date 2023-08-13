package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestHandle(t *testing.T) {
	cases := []struct {
		args          []string
		tailArgs      []string
		opts          CLIOptions
		expectedError error
	}{
		{
			args:          []string{},
			tailArgs:      []string{},
			opts:          CLIOptions{},
			expectedError: nil,
		},
		{
			args:          []string{"help"},
			tailArgs:      []string{"help"},
			opts:          CLIOptions{},
			expectedError: nil,
		},
		{
			args:          []string{"test.txt"},
			tailArgs:      []string{"test.txt"},
			opts:          CLIOptions{},
			expectedError: nil,
		},
		{
			args:          []string{"test.txt", "sub"},
			tailArgs:      []string{"test.txt", "sub"},
			opts:          CLIOptions{},
			expectedError: nil,
		},
		{
			args:          []string{"-b", "test.txt", "sub"},
			tailArgs:      []string{"test.txt", "sub"},
			opts:          CLIOptions{},
			expectedError: nil,
		},
		{
			args:          []string{"-b", "10", "test.txt"},
			tailArgs:      []string{"test.txt"},
			opts:          CLIOptions{},
			expectedError: nil,
		},
		{
			args:          []string{"-b", "test.txt"},
			tailArgs:      []string{"test.txt"},
			opts:          CLIOptions{},
			expectedError: nil,
		},
		{
			args:          []string{"-n", "10", "test.txt"},
			tailArgs:      []string{"test.txt"},
			opts:          CLIOptions{},
			expectedError: nil,
		},
		{
			args:          []string{"-n", "0", "test.txt"},
			tailArgs:      []string{"test.txt"},
			opts:          CLIOptions{},
			expectedError: fmt.Errorf("split: 0: illegal line count"),
		},
		{
			args:          []string{"-l", "10", "test.txt"},
			tailArgs:      []string{"test.txt"},
			opts:          CLIOptions{},
			expectedError: nil,
		},
		{
			args:          []string{"-l", "0", "test.txt"},
			tailArgs:      []string{"test.txt"},
			opts:          CLIOptions{},
			expectedError: fmt.Errorf("split: 0: illegal line count"),
		},
	}

	for _, c := range cases {
		t.Run(strings.Join(c.args, "_"), func(t *testing.T) {
			err := c.opts.Handle(c.args, c.tailArgs)
			if err != nil && c.expectedError == nil {
				t.Errorf("Unexpected error: %v", err)
			} else if err == nil && c.expectedError != nil {
				t.Errorf("Expected error, but got none")
			} else if err != nil && c.expectedError != nil && err.Error() != c.expectedError.Error() {
				t.Errorf("Expected error: %v, but got: %v", c.expectedError, err)
			}
		})
	}
}
