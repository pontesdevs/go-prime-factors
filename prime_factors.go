package go_prime_factors

func PrimeFactors(n int) []int {
	result := []int{}
	if n > 1 {
		for n%2 == 0 {
			result = append(result, 2)
			n /= 2
		}
	}
	if n > 1 {
		result = append(result, n)
	}
	return result
}
