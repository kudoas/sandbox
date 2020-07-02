package main

import "fmt"

// pointer 値のメモリアドレス
func pointers() {
	i, j := 42, 2701
	p := &i            // オペランドへのポインタを示す
	*p = 21            // set i through the pointer
	fmt.Println(*p, j) // * ポインタの指す先の変数を示す
}

func main() {
	pointers()
}
