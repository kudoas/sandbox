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

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

// Exercise: Loops and Functions
func Sqrt(x float64) float64 {
	z := float64(1)

	for i := 0; i < 10; i++ {
		diff := (z*z - x) / (2 * z)
		if diff == 0 {
			break
		}
		z -= diff
	}

	return z
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

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	fmt.Println(Sqrt(4))
}
