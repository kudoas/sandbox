package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
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

func say_hello() {
	defer fmt.Println("World")

	fmt.Println("Hello")
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

	// 選択された case だけを実行してそれに続く全ての case は実行されません
	// break ステートメントが Go では自動的に提供されます。
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s. \n", os)
	}

	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	say_hello()

	// LIFO: https://ja.wikipedia.org/wiki/LIFO
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}
