package main

import "fmt"

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
}
