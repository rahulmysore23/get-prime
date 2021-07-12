package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rahulmysore23/get-prime/pkg/logger"
	"github.com/rahulmysore23/get-prime/pkg/prime"
)

type PrimeSvr struct {
	Primes prime.Prime
	Log    logger.Logger
}

type GetPrimeOut struct {
	Intime   time.Time     `json:"in_time"`
	Outtime  time.Time     `json:"out_time"`
	Latency  time.Duration `json:"latency"`
	IsPrime  bool          `json:"is_prime"`
	PrimeNum int64         `json:"prime_num"`
}

type CheckPrimeOut struct {
	Intime  time.Time     `json:"in_time"`
	Outtime time.Time     `json:"out_time"`
	Latency time.Duration `json:"latency"`
	IsPrime bool          `json:"is_prime"`
}

type ErrorOut struct {
	Error   string        `json:"error"`
	Intime  time.Time     `json:"in_time"`
	Outtime time.Time     `json:"out_time"`
	Latency time.Duration `json:"latency"`
}

func (p PrimeSvr) GetPrime(c *gin.Context) {
	startTime := time.Now()

	num, err := strconv.ParseInt(c.Param("num"), 10, 64)
	if err != nil {
		// return failed response
		p.AbortRequestWithError(c, err, startTime)
		return
	}

	primeNum, isPrime := p.Primes.GetPrime(num)
	latency := time.Since(startTime).Round(time.Nanosecond)

	out := GetPrimeOut{
		Intime:   startTime,
		Outtime:  time.Now(),
		IsPrime:  isPrime,
		PrimeNum: primeNum,
		Latency:  latency,
	}

	c.JSON(http.StatusOK, out)

}

func (p PrimeSvr) CheckPrime(c *gin.Context) {
	startTime := time.Now()

	num, err := strconv.ParseInt(c.Param("num"), 10, 64)
	if err != nil {
		// return failed response
		p.AbortRequestWithError(c, err, startTime)
		return
	}

	isPrime := p.Primes.CheckPrime(num)
	latency := time.Since(startTime).Round(time.Nanosecond)

	out := CheckPrimeOut{
		Intime:  startTime,
		Outtime: time.Now(),
		IsPrime: isPrime,
		Latency: latency,
	}

	c.JSON(http.StatusOK, out)
}

func NewPrimeSvr(primes prime.Prime, log logger.Logger) PrimeSvr {
	return PrimeSvr{
		Primes: primes,
		Log:    log,
	}
}

func (p PrimeSvr) AbortRequestWithError(c *gin.Context, err error, startTime time.Time) {

	p.Log.LogError(err.Error())

	latency := time.Since(startTime).Round(time.Nanosecond)

	errOut := ErrorOut{
		Latency: latency,
		Intime:  startTime,
		Outtime: time.Now(),
		Error:   err.Error(),
	}

	c.JSON(http.StatusBadRequest, errOut)
}
