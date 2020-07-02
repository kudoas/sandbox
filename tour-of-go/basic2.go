package main

import (
	"fmt"
	"math"
)

func for1() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

func infiniteLoop() {
	for {
	}
}

func pow(x, n, lim float64) float64 {
	// if v := math.Pow(x, n); v < lim {
	// 	return v
	// } else {
	// 	fmt.Printf("%g >= %g\n", v, lim)
	// }
	// // can't use v here, though
	// return lim

	// recommend!!
	v := math.Pow(x, n)
	if v < lim {
		return v
	}
	fmt.Printf("%g >= %g\n", v, lim)
	return lim
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}
