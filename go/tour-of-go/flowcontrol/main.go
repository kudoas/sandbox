package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += 1
	}
	for sum < 100 {
		sum += sum
	}

	fmt.Println(sum)

	// infinite loop
	// for {
	// }

	fmt.Println(sqrt(2), sqrt(-4))
}
