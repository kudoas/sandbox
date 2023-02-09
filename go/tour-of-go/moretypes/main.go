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
}
