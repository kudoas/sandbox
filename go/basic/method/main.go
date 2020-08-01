package main

import "fmt"

type Calc struct {
	val1, val2 int
}

type Calc2 struct {
	val1, val2, val3 int
}

func Add(c Calc) int {
	return c.val1 + c.val2
}

// func (receiver, type) 関数名 (引数) 戻り値の型 {}
// 同じ変数名が付けられる、影響範囲が分かる
func (c Calc) Add() int {
	return c.val1 + c.val2
}

func (c Calc2) Add() int {
	return c.val1 + c.val2 + c.val3
}

func main() {
	calc := Calc{2, 3}
	calc2 := Calc2{1, 2, 3}
	fmt.Println(Add(calc))
	fmt.Println(calc.Add())
	fmt.Println(calc2.Add())
	fmt.Println(calc2.val1)
}
