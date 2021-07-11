package brute

import "math"

type Brute struct {
}

// CheckPrime returns true if the input number is prime
func (b Brute) CheckPrime(number int64) bool {
	for i := int64(2); i <= int64(math.Floor(float64(number)/2)); i++ {
		if number%i == 0 {
			return false
		}
	}
	return number > 1
}

// GetPrime returns the prime number which is highest number less than the input number
func (b Brute) GetPrime(number int64) (int64, bool) {

	var prime int64

	for p := number - 1; p > 1; p-- {
		if b.CheckPrime(p) {
			prime = p
			break
		}
	}

	return prime, b.CheckPrime(number)
}

// NewBrute creates
func NewBrute() Brute {
	return Brute{}
}
