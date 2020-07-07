package basic

import (
	"fmt"
	"math"
)

// ポインタ 値のメモリアドレス
func Pointers() {
	// var p *int
	// fmt.Println(p)
	i, j := 42, 32
	p := &i // & iへのポインタを探すオペランド
	// 0xc0000140c8 42 32
	fmt.Println(p, i, j)
	// * ポインタの指す先の変数(つまりポインタに対して使うオペランド)
	fmt.Println(*p)

	p = &j
	// dereferencing
	*p = *p * 64
	fmt.Println(j)
	i = i * 62
	fmt.Println(i)
}

// Vertex struct
type Vertex struct {
	X int
	Y int
}

// structのフィールドへは.を使ってアクセスする
func StructField() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v)

	// pointerを使ってアクセスもできる
	p := &v // 0xc00006cf28
	p.X = 1e9
	fmt.Println(v)
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

// array, sliceの違い！！
// array 固定長
func Arrays() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

// slice 可変長
func Slice() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes[1:4]
	// reference
	primes[2] = 37
	fmt.Println(primes, s)

	ss := []int{2, 3, 5, 7, 11, 13}
	// Slice the slice to give it zero length.
	ss = ss[:0]
	printSlice(ss)

	// Extend its length.
	ss = ss[:4]
	printSlice(ss)

	// Drop its first two values.
	ss = ss[2:]
	printSlice(ss)

	var sl []int
	printSlice(sl)
	if sl == nil {
		fmt.Println("nil")
	}
}

func MakeSlice() {
	// create a slice with make
	a := make([]int, 5)
	printSlice(a)
	b := make([]int, 0, 10)
	printSlice(b)
	c := b[:2]
	printSlice(c)
	d := c[2:5]
	printSlice(d)

	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	var appendSlice []int
	printSlice(appendSlice)
	appendSlice = append(appendSlice, 0)
	printSlice(appendSlice)
	appendSlice = append(appendSlice, 2, 3, 4)
	printSlice(appendSlice)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

// range
func RangeFunc() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	fmt.Println(pow)
	for index, value := range pow {
		fmt.Printf("2**%d = %d\n", index, value)
	}
	for _, value := range pow {
		fmt.Println(value)
	}
}

// Vertex1 hoge
type Vertex1 struct {
	Lat, Long float64
}

func mapFunc() {
	m := make(map[string]Vertex1)
	// map[Bell Labs:{40.68433 -74.39967}]
	m["Bell Labs"] = Vertex1{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"], m)
}

// map literals
func mapLiterals() map[string]Vertex1 {
	var m = map[string]Vertex1{
		"Bell Labs": Vertex1{
			40.68433, -74.39967,
		},
		"Google": {
			37.42202, -122.08408,
		},
	}

	// update, create
	m["Facevook"] = Vertex1{
		4.4444, -4.4444,
	}
	// delete
	delete(m, "Facevook")

	// is exit
	_, ok := m["Answer"]
	fmt.Println(ok)

	return m
}

// fuction value
func compute(f func(float64, float64) float64) float64 {
	return f(3, 4)
}

func Computer() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(compute(hypot))

}

// function closures
func Outer() func(int) int {
	sum := 0
	inner := func(x int) int {
		sum += x
		return sum
	}
	return inner
}

func Closure() {
	pos, neg := Outer(), Outer()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
