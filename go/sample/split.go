package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func splitByBytes(path string, bytesPerFile int) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	if bytesPerFile == 0 {
		outputFile, err := os.Create("1")
		if err != nil {
			return err
		}
		_, err = io.Copy(outputFile, file)
		if err != nil {
			return err
		}
		return nil
	}

	buffer := make([]byte, bytesPerFile)
	r := bufio.NewReader(file)

	for i := 1; ; i++ {
		n, err := r.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		outputFileName := fmt.Sprintf("%d", i)
		outputFile, err := os.Create(outputFileName)
		if err != nil {
			return err
		}
		outputFile.Write(buffer[:n])
		err = outputFile.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

func splitByLines(file *os.File, linesPerFile int) error {
	scanner := bufio.NewScanner(file)
	var lines []string

	index := 1
	for scanner.Scan() {
		lines = append(lines, scanner.Text())

		if len(lines) >= linesPerFile {
			outputFileName := fmt.Sprintf("%d", index)
			outputFile, err := os.Create(outputFileName)
			if err != nil {
				return err
			}
			for _, line := range lines {
				outputFile.WriteString(line + "\n")
			}
			err = outputFile.Close()
			if err != nil {
				return err
			}
			lines = nil
			index++
		}
	}

	if len(lines) > 0 {
		outputFileName := fmt.Sprintf("%d", index)
		outputFile, err := os.Create(outputFileName)
		if err != nil {
			return err
		}
		for _, line := range lines {
			outputFile.WriteString(line + "\n")
		}
		err = outputFile.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

func splitByChunks(file *os.File, chunksPerFile int) error {
	chunkSize := fileSize(file) / int64(chunksPerFile)
	buffer := make([]byte, chunkSize)
	r := bufio.NewReader(file)

	for i := 1; i <= chunksPerFile; i++ {
		outputFileName := fmt.Sprintf("%d", i)
		outputFile, err := os.Create(outputFileName)
		if err != nil {
			return err
		}
		_, err = r.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		_, err = outputFile.Write(buffer)
		if err != nil {
			return err
		}
		err = outputFile.Close()
		if err != nil {
			return err
		}
	}

	return nil
}

func fileSize(file *os.File) int64 {
	fileInfo, _ := file.Stat()
	return fileInfo.Size()
}
