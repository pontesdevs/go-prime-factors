package go_prime_factors

func PrimeFactors(n int) []int {
	result := []int{}
	if n > 1 {
		for divisor := 2; divisor <= n; divisor++ {
			for n%divisor == 0 {
				result = append(result, divisor)
				n /= divisor
			}
		}
	}
	return result
}
