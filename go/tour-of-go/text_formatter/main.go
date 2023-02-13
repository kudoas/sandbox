package main

import (
	"bufio"
	"fmt"
	"os"
)

// command
// $ cat text_formatter/sample.txt | go run text_formatter/main.go
func main() {
	not_last := true
	scanner := bufio.NewScanner(os.Stdin)

	for not_last {
		not_last = scanner.Scan()
		t := scanner.Text()
		fmt.Printf("%s\n", t)
	}
}
