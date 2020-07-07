package basic

import (
	"fmt"
	"math"
	"runtime"
)

func For1() {
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

func Pow(x, n, lim float64) float64 {
	// if v := math.Pow(x, n); v < lim {
	// 	return v
	// } else {
	// 	fmt.Printf("%g >= %g\n", v, lim)
	// }
	// // can't use v here, though
	// return lim

	// golint マジ神！
	v := math.Pow(x, n)
	if v < lim {
		return v
	}
	fmt.Printf("%g >= %g\n", v, lim)
	return lim
}

func Sqrt(x float64) string {
	if x < 0 {
		return Sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func Switch1() {
	switch os := runtime.GOOS; os {
	// case 条件が一致すれば自動でbreak
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.", os)
	}
}

// deferへ渡した関数の実行を、呼び出し元の関数の終わり(returnする)まで遅延させる
// 評価はされるが最後に実行する
func Defer1() {
	defer fmt.Println("world")
	fmt.Println("hello")
}

func DeferLikeStack() {
	fmt.Println("counting")
	defer fmt.Println("done")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
