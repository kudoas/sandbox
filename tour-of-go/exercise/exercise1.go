package exercise

import (
	"math"
)

// Exercise: Loops and Functions

// Sqrt1 大文字の時はコメントをつけるべし！
func Sqrt1(x float64, count int) float64 {
	z := float64(1)
	for i := 0; i < count; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

// Sqrt2 誤差がめっちゃ小さくなるまで
func Sqrt2(x float64) float64 {
	preZ, z := x, 1.0
	for math.Abs(z-preZ) >= 1.0e-6 {
		preZ, z = z, z-(z*z-x)/(2*z)
	}
	return z
}
