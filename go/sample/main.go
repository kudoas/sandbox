package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	linePtr := flag.Int("l", 0, "-l: Create split files line_count lines in length.")
	chunkPtr := flag.Int("n", 0, "-n: Split file into chunk_count smaller files.")
	bytePtr := flag.Int("b", 0, "-b: Create split files byte_count bytes in length.")
	flag.Parse()

	cliOptions := &CLIOptions{ByteCount: *bytePtr, LineCount: *linePtr, ChunkCount: *chunkPtr}
	tailArgs := flag.Args()
	err := cliOptions.Parse(os.Args[1:], tailArgs)
	if err != nil {
		fmt.Println(err)
	}
}

type CLIOptions struct {
	ByteCount  int
	LineCount  int
	ChunkCount int
}

func (opts *CLIOptions) Parse(args []string, tailArgs []string) error {
	if len(args) == 0 || len(tailArgs) != 1 || len(args) > 0 && args[0] == "help" {
		usage := `usage:
split [-l line_count] [file [prefix]]
split -b byte_count [file [prefix]]
split -n chunk_count [file [prefix]]`
		fmt.Print(usage)
		return nil
	}

	switch args[0] {
	case "-b":
		splitByBytes(tailArgs[0], opts.ByteCount)
	case "-l":
		arg, err := strconv.Atoi(args[1])
		if err != nil {
			return err
		}
		if arg <= 0 {
			return fmt.Errorf("split: %d: illegal line count", arg)
		}

		file, err := os.Open(tailArgs[0])
		if err != nil {
			return err
		}
		splitByLines(file, opts.LineCount)
	case "-n":
		arg, err := strconv.Atoi(args[1])
		if err != nil {
			return err
		}
		if arg <= 0 {
			return fmt.Errorf("split: %d: illegal line count", arg)
		}
		file, err := os.Open(tailArgs[0])
		if err != nil {
			return err
		}
		splitByChunks(file, opts.ChunkCount)
	default:
		file, err := os.Open(tailArgs[0])
		if err != nil {
			return err
		}
		splitByLines(file, 1000)
	}
	return nil
}

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
