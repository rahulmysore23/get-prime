package seive

import "math"

type Seive struct {
	Primes []bool
}

// CheckPrime returns true if the input number is prime
func (s Seive) CheckPrime(number int64) bool {
	return s.Primes[number]
}

// GetPrime returns the prime number which is highest number less than the input number
func (s Seive) GetPrime(number int64) (int64, bool) {
	var prime int64

	// TODO - use binary search
	for p := number - 1; p > 1; p-- {
		if s.Primes[p] {
			prime = p
			break
		}
	}

	return prime, s.Primes[number]
}

// NewSeive creates a seive with size n
func NewSeive(n int64) Seive {

	primes := make([]bool, n)

	for i := int64(2); i < n; i++ {
		primes[i] = true
	}

	for p := int64(2); p <= int64(math.Floor(float64(n)/2)); p++ {
		if primes[p] {
			for i := p * 2; i <= n; i += p {
				primes[i] = false
			}
		}
	}

	return Seive{
		Primes: primes,
	}

}
