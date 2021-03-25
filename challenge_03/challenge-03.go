package challenge_03

func Sum(n int) int {
	total := 0

	for i := 1; i <= n; i++ {
		total += i
	}

	return total
}
