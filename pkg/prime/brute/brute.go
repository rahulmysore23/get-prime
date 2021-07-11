package prime

type Brute struct {
}

func (b Brute) CheckPrime(number int64) bool {
	return false
}

func (b Brute) GetPrime(number int64) int64 {
	return number
}
