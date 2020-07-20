package exercise

// Fibonacci closure
func Fibonacci() func() []int {
	squence := make([]int, 2)
	squence[0] = 0
	squence[1] = 1
	now := 0
	formula := func() []int {
		next := squence[now] + squence[now+1]
		squence = append(squence, next)
		now++
		return squence
	}
	return formula
}
