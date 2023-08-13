package main

import (
	"os"
	"testing"
)

func TestSplitByBytes(t *testing.T) {
	cases := []struct {
		content       string
		bytesPerFile  int
		expectedFiles []string
		expectedError bool
	}{
		{
			content:       "file: larger than bytesPerFile",
			bytesPerFile:  10,
			expectedFiles: []string{"1", "2", "3"},
			expectedError: false,
		},
		{
			content:       "file: smaller than bytesPerFile",
			bytesPerFile:  100000,
			expectedFiles: []string{"1"},
			expectedError: false,
		},
		{
			content:       "0 byte",
			bytesPerFile:  0,
			expectedFiles: []string{"1"},
			expectedError: false,
		},
	}

	for _, c := range cases {
		t.Run(c.content, func(t *testing.T) {
			inputPath := "test_input.txt"
			err := createTestFile(t, inputPath, c.content)
			if err != nil {
				t.Fatalf("Failed to create test input file: %v", err)
			}
			defer os.Remove(inputPath)

			err = splitByBytes(inputPath, c.bytesPerFile)
			if err != nil && !c.expectedError {
				t.Fatalf("splitByBytes returned an unexpected error: %v", err)
			} else if err == nil && c.expectedError {
				t.Fatalf("Expected error, but got none")
			}
			for _, expectedFile := range c.expectedFiles {
				_, err := os.Stat(expectedFile)
				if err != nil {
					t.Errorf("File %s was not created as expected: %v", expectedFile, err)
				}
			}
			verifyFileSize(t, c.bytesPerFile, c.expectedFiles)
			cleanTestFile(t, c.expectedFiles)
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

func verifyFileSize(t *testing.T, expectedFileSize int, fileNames []string) {
	t.Helper()
	for _, fileName := range fileNames {
		file, err := os.Open(fileName)
		if err != nil {
			t.Errorf("Error opening file %s: %v", fileName, err)
			continue
		}
		defer file.Close()

		info, err := file.Stat()
		if err != nil {
			t.Error(err)
		}
		if expectedFileSize != 0 && int(info.Size()) > expectedFileSize {
			t.Errorf("File %s size mismatch, Expected: %d, but got: %d", fileName, expectedFileSize, info.Size())
		}
	}
}

func cleanTestFile(t *testing.T, fileNames []string) {
	t.Helper()
	for _, fileName := range fileNames {
		if err := os.Remove(fileName); err != nil {
			t.Error(err)
		}
	}
}
