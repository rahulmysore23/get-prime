package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rahulmysore23/get-prime/pkg/prime"
)

type PrimeSvr struct {
	Primes prime.Prime
}

func (p PrimeSvr) GetPrime(c *gin.Context) {
	num, err := strconv.ParseInt(c.Param("num"), 10, 64)
	if err != nil {
		// return failed response
		panic(err) // TODO - temp useage, remove this
	}

	primeNum, isPrime := p.Primes.GetPrime(num)

	c.JSON(http.StatusOK, map[string]interface{}{
		"in_time":   "",
		"out_time":  "",
		"latency":   "",
		"is_prime":  isPrime,
		"prime_num": primeNum,
	})

}

func (p PrimeSvr) CheckPrime(c *gin.Context) {
	num, err := strconv.ParseInt(c.Param("num"), 10, 64)
	if err != nil {
		// return failed response
		panic(err) // TODO - temp useage, remove this
	}

	isPrime := p.Primes.CheckPrime(num)

	c.JSON(http.StatusOK, map[string]interface{}{
		"in_time":  "",
		"out_time": "",
		"latency":  "",
		"is_prime": isPrime,
	})
}
