package emath

/**
 * Get fibonacci n-length
 */
func Calculate(n int) []int {
	return calculate(n)
}

/**
 * Get nth number from fibonacci
 */
func CalculateOne(n int) int {
	f := calculate(n)
	return f[n-1]
}

func calculate(n int) []int {
	fibonacci := make([]int, n)
	for i := 1; i < n; i++ {
		if i == 1 {
			fibonacci[i] = i
		} else {
			fibonacci[i] = fibonacci[i-1] + fibonacci[i-2]
		}
	}

	return fibonacci
}