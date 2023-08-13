package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestSplitByBytes(t *testing.T) {
	testCases := []struct {
		name           string
		inputContent   string
		bytesPerFile   int
		expectedFiles  []string
		expectedErrors bool
	}{
		{
			name:           "SplitByBytes with non-zero bytesPerFile",
			bytesPerFile:   10,
			expectedFiles:  []string{"1", "2", "3"},
			expectedErrors: false,
		},
		{
			name:           "SplitByBytes with zero bytesPerFile",
			bytesPerFile:   0,
			expectedFiles:  []string{"1"},
			expectedErrors: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			inputPath := "test_input.txt"
			err := createTestFile(inputPath, "This is a test input.")
			if err != nil {
				t.Fatalf("Failed to create test input file: %v", err)
			}
			defer os.Remove(inputPath)

			err = splitByBytes(inputPath, tc.bytesPerFile)
			if err != nil && !tc.expectedErrors {
				t.Fatalf("splitByBytes returned an unexpected error: %v", err)
			} else if err == nil && tc.expectedErrors {
				t.Fatalf("Expected error, but got none")
			}

			for _, expectedFile := range tc.expectedFiles {
				_, err := os.Stat(expectedFile)
				if err != nil {
					t.Errorf("File %s was not created as expected: %v", expectedFile, err)
				}
			}

			verifyFileContent(t, tc.inputContent, tc.expectedFiles)
		})
	}
}

func createTestFile(t *testing.T, filePath, content string) error {
	t.Helper()
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}

func verifyFileContent(t *testing.T, expectedContent string, fileNames []string) {
	t.Helper()
	for _, fileName := range fileNames {
		file, err := os.Open(fileName)
		if err != nil {
			t.Errorf("Error opening file %s: %v", fileName, err)
			continue
		}
		defer file.Close()

		var contentBuffer bytes.Buffer
		_, err = io.Copy(&contentBuffer, file)
		if err != nil {
			t.Errorf("Error reading file %s: %v", fileName, err)
			continue
		}

		if contentBuffer.String() != expectedContent {
			t.Errorf("File %s content mismatch. Expected: %s, Got: %s", fileName, expectedContent, contentBuffer.String())
		}
	}
}
