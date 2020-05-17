package count_primes

func countPrimes(n int) int {
	var ret int
	for i := 2; i < n; i++ {
		if isPrime(i) {
			ret++
		}
	}
	return ret
}

func isPrime(value int) bool {
	if value <= 3 {
		return value >= 2
	}
	if value%2 == 0 || value%3 == 0 {
		return false
	}
	for i := 5; i*i <= value; i += 6 {
		if value%i == 0 || value%(i+2) == 0 {
			return false
		}
	}
	return true
}

func countPrimes1(n int) int {
	isNotPrime := make([]bool, n)
	for i := 2; i*i < n; i++ {
		if isNotPrime[i] {
			continue
		}
		for j := i * i; j < n; j += i {
			isNotPrime[j] = true
		}
	}
	var ret int
	for i := 2; i < n; i++ {
		if !isNotPrime[i] {
			ret++
		}
	}

	return ret
}
