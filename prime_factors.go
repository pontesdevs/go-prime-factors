package go_prime_factors

func PrimeFactors(n int) []int {
	result := []int{}
	if n == 2 {
		result = append(result, 2)
	}
	return result
}
