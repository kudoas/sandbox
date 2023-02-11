package exercise

import (
	"golang.org/x/tour/pic"
)

// Pic boardを作って関数で埋める
func Pic(dx, dy int) [][]uint8 {
	board := make([][]uint8, dx)
	for i := 0; i < dx; i++ {
		board[i] = make([]uint8, dy)
	}
	for i := 0; i < dy; i++ {
		for j := 0; j < dx; j++ {
			board[i][j] = Calculation(i, j)
		}
	}
	return board
}

func Calculation(i, j int) uint8 {
	return uint8(i * j)
}

func DrawGraph() {
	pic.Show(Pic)
}
