package main

import (
	"fmt"
)

type DB struct {
	When int
	What string
}

// method
func (db *DB) Sample(a *[]int) {
	*a = append(*a, 13)
}

func main() {
	a := make([]int, 0)
	var db *DB
	// 構造体に紐づいたmethodを呼び出せる！
	if db.Sample(&a); len(a) < 2 {
		fmt.Println("hgoe")
	}
	fmt.Println(a)
}
