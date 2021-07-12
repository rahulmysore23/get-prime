package http_server

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rahulmysore23/get-prime/config"
	"github.com/rahulmysore23/get-prime/http_server/handlers"
	"github.com/rahulmysore23/get-prime/pkg/logger"
	"github.com/rahulmysore23/get-prime/pkg/logger/logrus"
	"github.com/rahulmysore23/get-prime/pkg/prime"
	"github.com/rahulmysore23/get-prime/pkg/prime/brute"
	"github.com/rahulmysore23/get-prime/pkg/prime/seive"
	"github.com/rahulmysore23/get-prime/pkg/utilities"
	"github.com/urfave/cli/v2"
)

var (
	primes prime.Prime
	log    logger.Logger
	flags  config.AppFlags
)

// Start will use a serve config pkg to setup the server
func GetCliApp() *cli.App {
	app := cli.NewApp()

	// Set flags for app - i.e, Env variables
	app.Flags = config.GetFlags(&flags)
	app.Action = Run
	return app
}

// Run will setup the server depdendencies
func Run(_ *cli.Context) error {

	if flags.PrimeAlgorithm == utilities.BRUTE {
		primes = brute.NewBrute()
	} else if flags.PrimeAlgorithm == utilities.SEIVE {
		primes = seive.NewSeive(flags.MaxPrimes)
	} else {
		return fmt.Errorf("wrong algo mentioned in env")
	}

	if flags.PrimeLogger {
		primes = prime.LogPrime{Prime: primes}
	}

	return InitSever(primes)
}

// InitSever will inject resources and create REST endpoints
func InitSever(primes prime.Prime) error {
	log = logrus.NewLogrus(flags.LogLevel)

	routes := gin.New()
	routes.Use(
		CORS(),
		gin.Recovery(),
		gin.Logger(),
	)

	v1 := routes.Group("/v1")

	svr := handlers.NewPrimeSvr(primes, log)

	v1.GET("/get_prime/:num", svr.GetPrime)
	v1.GET("/check_prime/:num", svr.CheckPrime)

	// For test - REMOVE AFTER ADDING UI
	routes.GET("/", func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "text/html")
		c.Writer.WriteHeader(http.StatusOK)
		io.WriteString(c.Writer, utilities.DefaultHtml)
	})

	fmt.Printf("webserver listening on port %v\n", "6060") // TODO - Move to env
	err := http.ListenAndServe(":6060", routes)
	if err != nil {
		log.LogError(err.Error())
		return err
	}

	return nil
}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.Writer.Header()
		header.Set("Access-Control-Allow-Origin", "*")
		header.Set("Access-Control-Allow-Methods", "*")
		header.Set("Access-Control-Allow-Headers", "Content-Type,Authorization")
		header.Set("Access-Control-Max-Age", "86400")

		if c.Request.Method == http.MethodOptions {
			c.Writer.WriteHeader(http.StatusOK)
			return
		}

		c.Next()
	}
}
