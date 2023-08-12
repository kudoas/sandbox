package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestHandle(t *testing.T) {
	handleTest := []struct {
		args        []string
		tailArgs    []string
		opts        CLIOptions
		expectedErr error
	}{
		{
			args:        []string{},
			tailArgs:    []string{},
			opts:        CLIOptions{},
			expectedErr: nil,
		},
		{
			args:        []string{"help"},
			tailArgs:    []string{"help"},
			opts:        CLIOptions{},
			expectedErr: nil,
		},
		{
			args:        []string{"test.txt"},
			tailArgs:    []string{"test.txt"},
			opts:        CLIOptions{},
			expectedErr: nil,
		},
		{
			args:        []string{"test.txt", "sub"},
			tailArgs:    []string{"test.txt", "sub"},
			opts:        CLIOptions{},
			expectedErr: nil,
		},
		{
			args:        []string{"-b", "test.txt", "sub"},
			tailArgs:    []string{"test.txt", "sub"},
			opts:        CLIOptions{},
			expectedErr: nil,
		},
		{
			args:        []string{"-b", "10", "test.txt"},
			tailArgs:    []string{"test.txt"},
			opts:        CLIOptions{},
			expectedErr: nil,
		},
		{
			args:        []string{"-b", "test.txt"},
			tailArgs:    []string{"test.txt"},
			opts:        CLIOptions{},
			expectedErr: nil,
		},
		{
			args:        []string{"-n", "10", "test.txt"},
			tailArgs:    []string{"test.txt"},
			opts:        CLIOptions{},
			expectedErr: nil,
		},
		{
			args:        []string{"-n", "0", "test.txt"},
			tailArgs:    []string{"test.txt"},
			opts:        CLIOptions{},
			expectedErr: fmt.Errorf("split: 0: illegal line count"),
		},
		{
			args:        []string{"-l", "10", "test.txt"},
			tailArgs:    []string{"test.txt"},
			opts:        CLIOptions{},
			expectedErr: nil,
		},
		{
			args:        []string{"-l", "0", "test.txt"},
			tailArgs:    []string{"test.txt"},
			opts:        CLIOptions{},
			expectedErr: fmt.Errorf("split: 0: illegal line count"),
		},
	}

	for _, tt := range handleTest {
		t.Run(strings.Join(tt.args, "_"), func(t *testing.T) {
			err := tt.opts.Handle(tt.args, tt.tailArgs)
			if err != nil && tt.expectedErr == nil {
				t.Errorf("Unexpected error: %v", err)
			} else if err == nil && tt.expectedErr != nil {
				t.Errorf("Expected error, but got none")
			} else if err != nil && tt.expectedErr != nil && err.Error() != tt.expectedErr.Error() {
				t.Errorf("Expected error: %v, but got: %v", tt.expectedErr, err)
			}
		})
	}
}
