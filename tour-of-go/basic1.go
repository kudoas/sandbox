package main

import (
	"fmt"
	"math"
)

// Create a huge number by shifting a 1 bit left 100 places.
// In other words, the binary number that is 1 followed by 100 zeroes.
// Shift it right again 99 places, so we end up with 1<<1, or 2.
const (
	Big   = 1 << 100
	Small = Big >> 99
)

// int は64-bitの整数を保持できますが、それでは足りないことが時々あります。そういったときにconstを活用しましょう)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

// func main() {
// 	fmt.Println(needInt(Small))
// 	fmt.Println(needFloat(Small))
// 	fmt.Println(needFloat(Big))
// }

// type suggestion
func main2() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"
	fmt.Println(i, j, k, c, python, java)

	// type conversions
	var x, y int = 3, 4
	// 明示的な変換が必要
	var f float64 = math.Sqrt(float64(x*x + y*y))
	// var f := math.Sqrt(float64(x*x + y*y))
	// var f float64 = math.Sqrt(x*x + y*y) error!
	var z uint = uint(f)
	fmt.Println(x, y, z)

	// type inferenct
	// v := 0.12 float

	// constant
	const world = "sekai" // constant

}

// 基本型
// bool
// string
// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64 uintptr
// byte // uint8 の別名
// rune // int32 の別名
//      // Unicode のコードポイントを表す
// float32 float64
// complex64 complex128

// uint: 32bit number
// complex64: 複素数
