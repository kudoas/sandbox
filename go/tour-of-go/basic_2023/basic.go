// main package で実行する
package main

import (
	"fmt"
	"math"
)

func main() {
	// 最初の文字が大文字で始まる名前は、外部のパッケージから参照できるエクスポート（公開）された名前( exported name )
	fmt.Println(math.Pi)
}
