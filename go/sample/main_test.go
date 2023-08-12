package main

import (
	"fmt"
	"os"
	"testing"
)

func TestSplitByBytes(t *testing.T) {
	testCases := []struct {
		inputFileName string
		bytesPerFile  int
	}{
		{"memo.txt", 10},
		{"memo.txt", 20},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("bytesPerFile=%d", testCase.bytesPerFile), func(t *testing.T) {
			inputFile, err := os.Open(testCase.inputFileName)
			if err != nil {
				t.Errorf("Failed to open input file: %v", err)
				return
			}
			defer inputFile.Close()

			err = splitByBytes(inputFile, testCase.bytesPerFile)
			if err != nil {
				t.Errorf("Error during splitting: %v", err)
			}
		})
	}
}
