package main

import "fmt"

// pointer 値のメモリアドレス
func pointers() {
	i, j := 42, 2701
	p := &i            // オペランドへのポインタを示す
	*p = 21            // set i through the pointer
	fmt.Println(*p, j) // * ポインタの指す先の変数を示す
}

// Vertex struct
type Vertex struct {
	X int
	Y int
}

func outputVertex() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
}

func pointerToStruct() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

// array 固定長
func arrays() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

// slice 可変長
func slice() {
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
func rangeFunc() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
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

func main() {
	// pointers()
	// pointerToStruct()
	// fmt.Println(v1, p, v2, v3)
	// arrays()
	// slice()
	// rangeFunc()
	mapFunc()
}
