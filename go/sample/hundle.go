package main

import (
	"fmt"
	"os"
	"strconv"
)

type CLIOptions struct {
	ByteCount  int
	LineCount  int
	ChunkCount int
}

func (opts *CLIOptions) Handle(args []string, tailArgs []string) error {
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
