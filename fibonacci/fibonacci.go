package fibonacci

/**
 * Get fibonacci n-length
 */
func Calculate(n int) []int {
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

/**
 * Get nth number from fibonacci
 */
func CalculateOne(n int) int {
	f := Calculate(n)
	return f[n-1]
}