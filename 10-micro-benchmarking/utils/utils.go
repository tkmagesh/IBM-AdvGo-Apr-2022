package utils

import "math"

func GeneratePrimes(start, end int) []int {
	//primes := []int{}
	primes := make([]int, 0, 50)
	for no := start; no <= end; no++ {
		if IsPrime(no) {
			primes = append(primes, no)
		}
	}
	return primes
}

func IsPrime(no int) bool {
	for i := 2; i < no; i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}

func IsPrime_2(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}

func IsPrime_2_A(no int) bool {
	end := no / 2
	for i := 2; i <= end; i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}

func IsPrime_3(no int) bool {
	end := int(math.Sqrt(float64(no)))
	for i := 2; i <= end; i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
