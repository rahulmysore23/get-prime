package prime

type Seive struct {
	Primes []bool
}

func (s Seive) CheckPrime(number int64) bool {
	return false
}

func (s Seive) GetPrime(number int64) int64 {
	return number
}
