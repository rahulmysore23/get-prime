package prime

import (
	"fmt"
	"time"
)

type LogPrime struct {
	Prime Prime
}

func (l LogPrime) CheckPrime(number int64) bool {
	defer func(begin time.Time) {
		fmt.Printf("CheckPrime took %v", time.Since(begin).Round(time.Millisecond))
	}(time.Now())

	return l.Prime.CheckPrime(number)
}

func (l LogPrime) GetPrime(number int64) (int64, bool) {
	defer func(begin time.Time) {
		fmt.Printf("GetPrime took %v", time.Since(begin).Round(time.Millisecond))
	}(time.Now())

	return l.Prime.GetPrime(number)
}

func NewLogger(primes Prime) Prime {
	return LogPrime{
		Prime: primes,
	}
}
