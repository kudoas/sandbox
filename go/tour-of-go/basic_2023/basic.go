// main package で実行する
package main

import (
	"fmt"
	"math"
)

// MEMO: https://go.dev/blog/declaration-syntax
func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

// named return value
// リーダビリティが下がるので短い関数で使う
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	// 最初の文字が大文字で始まる名前は、外部のパッケージから参照できるエクスポート（公開）された名前( exported name )
	fmt.Println(math.Pi)

	fmt.Println(add(1, 2))
	fmt.Println(split((17)))

	// default は false
	// zero value
	var c, python, java bool
	var i int
	var elm string = "function"

	// 宣言代入
	kotlin := "mobile"

	fmt.Println(i, c, python, java, elm, kotlin)

	// 型変換（Type conversions）
	var x, y int = 3, 4
	var f = math.Sqrt(float64(x*x + y*y)) // 変換しないとタイプエラー
	fmt.Println(f)

	// 型類推
	v := "fmt" // change me!
	fmt.Printf("v is of type %T\n", v)

	// 定数
	// 文字(character)、文字列(string)、boolean、数値(numeric)のみ
	const world = 2.2
	fmt.Println(world)

}
