package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	linePtr := flag.Int("l", 0, "-l: Create split files line_count lines in length.")
	chunkPtr := flag.Int("n", 0, "-n: Split file into chunk_count smaller files.")
	bytePtr := flag.Int("b", 0, "-b: Create split files byte_count bytes in length.")
	flag.Parse()

	cliOptions := &CLIOptions{ByteCount: *bytePtr, LineCount: *linePtr, ChunkCount: *chunkPtr}
	err := cliOptions.Parse(os.Args[1:])
	if err != nil {
		fmt.Println(err)
	}
}

type CLIOptions struct {
	ByteCount  int ``
	LineCount  int ``
	ChunkCount int ``
}

func (opts *CLIOptions) Parse(args []string) error {
	// option の仕様
	// -b, -n, -l は同時に指定できない、どれか1つだけ
	// -b, -n, -l はいずれも0やマイナスを指定できない
	// option の validation をする層が必要そう

	switch args[0] {
	case "-b":
		file, err := os.Open(args[2])
		if err != nil {
			return err
		}
		splitByBytes(file, opts.ByteCount)
	case "-l":
		if args[1] == "0" {
			return fmt.Errorf("split: 0: illegal line count")
		}

		file, err := os.Open(args[2])
		if err != nil {
			return err
		}
		splitByLines(file, opts.LineCount)
	case "-n":
		if args[1] == "0" {
			return fmt.Errorf("split: 0: illegal number of chunks")
		}
		file, err := os.Open(args[2])
		if err != nil {
			return err
		}
		splitByChunks(file, opts.ChunkCount)
	default:
		file, err := os.Open(args[0])
		if err != nil {
			return err
		}
		splitByLines(file, 1000)
	}
	return nil
}

func splitByBytes(file *os.File, bytesPerFile int) error {
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
	for i := 1; ; i++ {
		n, err := file.Read(buffer)
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
		outputFile.Close()
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
			outputFile.Close()
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
		outputFile.Close()
	}

	return nil
}

func splitByChunks(file *os.File, chunksPerFile int) error {
	chunkSize := fileSize(file) / int64(chunksPerFile)
	buffer := make([]byte, chunkSize)

	for i := 1; i <= chunksPerFile; i++ {
		outputFileName := fmt.Sprintf("%d", i)
		outputFile, err := os.Create(outputFileName)
		if err != nil {
			return err
		}
		defer outputFile.Close()

		_, err = io.ReadFull(file, buffer)
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
	}

	return nil
}

func fileSize(file *os.File) int64 {
	fileInfo, _ := file.Stat()
	return fileInfo.Size()
}

// func validateOptions(opts *CLIOptions)

// func parse(args []string) (map[string]string, error) {
// 	var opts CLIOptions

// 	// 何も値が渡されてない場合は help を実行する
// 	if len(args) == 0 {
// 		return nil, fmt.Errorf("failed to parse")
// 	}

// 	options := make(map[string]string)
// 	filePath := args[0]

// 	for i := 1; i < len(args)-1; i += 2 {
// 		key := args[i]
// 		value := args[i+1]

// 		options[key] = value
// 	}

// 	// 指定されてないオプションが渡された場合

// 	// 最後にファイル名が渡されてない場合

// 	// OK

// 	fmt.Println(args)
// 	return args
// }

// option のバリデーションはこういう書き方ができそう！
// switch {
// case t.Hour() < 12:
// 	fmt.Println("Good morning!")
// case t.Hour() < 17:
// 	fmt.Println("Good afternoon.")
// default:
// 	fmt.Println("Good evening.")
// }

// ポインタレシーバを使う2つの理由があります。
// ひとつは、メソッドがレシーバが指す先の変数を変更するためです。
// ふたつに、メソッドの呼び出し毎に変数のコピーを避けるためです。 例えば、レシーバが大きな構造体である場合に効率的です。

// https://go-tour-jp.appspot.com/methods/21
// func main() {
// 	r := strings.NewReader("Hello, Reader!")

// 	b := make([]byte, 8)
// 	for {
// 		n, err := r.Read(b)
// 		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
// 		fmt.Printf("b[:n] = %q\n", b[:n])
// 		if err == io.EOF {
// 			break
// 		}
// 	}
// }
