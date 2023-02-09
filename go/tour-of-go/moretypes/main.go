package main

import "fmt"

type Vertex struct {
	X int
	Y int
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
}
