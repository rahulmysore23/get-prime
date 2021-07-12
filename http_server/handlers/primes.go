package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rahulmysore23/get-prime/pkg/prime"
)

type PrimeSvr struct {
	Primes prime.Prime
}

func (p PrimeSvr) GetPrime(c *gin.Context) {
	startTime := time.Now()

	num, err := strconv.ParseInt(c.Param("num"), 10, 64)
	if err != nil {
		// return failed response
		panic(err) // TODO - temp useage, remove this
	}

	primeNum, isPrime := p.Primes.GetPrime(num)

	var latency time.Duration

	latency = time.Since(startTime).Round(time.Millisecond)

	if latency == 0 {
		latency = time.Since(startTime).Round(time.Nanosecond)
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"in_time":   startTime,
		"out_time":  time.Now(),
		"latency":   latency,
		"is_prime":  isPrime,
		"prime_num": primeNum,
	})

}

func (p PrimeSvr) CheckPrime(c *gin.Context) {
	startTime := time.Now()

	num, err := strconv.ParseInt(c.Param("num"), 10, 64)
	if err != nil {
		// return failed response
		panic(err) // TODO - temp useage, remove this
	}

	isPrime := p.Primes.CheckPrime(num)

	var latency time.Duration

	latency = time.Since(startTime).Round(time.Millisecond)

	if latency == 0 {
		latency = time.Since(startTime).Round(time.Nanosecond)
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"in_time":  startTime,
		"out_time": time.Now(),
		"latency":  latency,
		"is_prime": isPrime,
	})
}

func NewPrimeSvr(primes prime.Prime) PrimeSvr {
	return PrimeSvr{
		Primes: primes,
	}
}
