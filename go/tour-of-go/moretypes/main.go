package main

import (
	"fmt"
	"strings"
)

type Vertex struct {
	X int
	Y int
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main() {
	i, j := 25, 28

	// & オペレータは、そのオペランド (operand) へのポインタを引き出します
	p := &i
	fmt.Println(p)
	fmt.Println(*p)

	// dereferencing or indirecting
	*p = 52
	fmt.Println(i)

	p = &j
	*p = *p / 14
	fmt.Println(j)

	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.Y, v.X)

	vv := Vertex{2, 4}
	pv := &vv
	pv.X = 1e9 // (*pv).X は 描くのだるいから省略できる
	fmt.Println(vv)

	// struct の初期値を割り当てられる
	var (
		v1 = Vertex{1, 2}  // has type Vertex
		v2 = Vertex{X: 1}  // Y:0 is implicit
		v3 = Vertex{}      // X:0 and Y:0
		vp = &Vertex{1, 2} // has type *Vertex
	)

	fmt.Println(v1, vp, v2, v3)

	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	var s []int = primes[1:4]
	fmt.Println(s)

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}

	e := names[1:3]
	e[0] = "XXX" // 参照渡し
	fmt.Println(names)

	// 長さのもたない配列リテラル
	q := []int{2, 3, 5, 7, 11, 13}
	r := []bool{true, false, true, true, false, true}
	sl := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(q, r, sl)

	// スライスの容量は、スライスの最初の要素から数えて、元となる配列の要素数
	// 途中から数えるとその分減る
	printSlice(s)
	printSlice(q)

	// Nil slices
	var ns []int
	fmt.Println(s, len(ns), cap(ns))
	if ns == nil {
		fmt.Println("nil")
	}

	// 0 が 5つ入る
	a1 := make([]int, 5)
	printSlice(a1)

	// 容量が 5 の slice
	b := make([]int, 0, 5)
	printSlice(b)

	// 容量は 5　のまま
	c := b[:2]
	printSlice(c)

	// 容量が 3　になる
	d := c[2:5]
	printSlice(d)

	// 2次元スライス string string
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	// https://pkg.go.dev/builtin#append
	s = append(s, 1)
	printSlice(s)
	s = append(s, 1, 2, 2)
	printSlice(s)
}
