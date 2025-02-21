package go_prime_factors

func PrimeFactors(n int) []int {
	result := []int{}
	if n > 1 {
		result = append(result, n)
	}
	return result
}
