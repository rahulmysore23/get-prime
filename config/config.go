package config

import "github.com/urfave/cli/v2"

const MAX_PRIME_SIZE = 10000

type AppFlags struct {
	Env            string
	PrimeAlgorithm string
	MaxPrimes      int64
	PrimeLogger    bool
	LogLevel       string
	Version        string
}

func GetFlags(flags *AppFlags) []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "Server-Environment",
			Value:       "LOCAL",
			EnvVars:     []string{"ENV"},
			Destination: &flags.Env,
		},
		&cli.StringFlag{
			Name:        "Prime-Number-Algorithm",
			Value:       "BRUTE",
			EnvVars:     []string{"PRIME_ALGORITHM"},
			Destination: &flags.PrimeAlgorithm,
		},
		&cli.Int64Flag{
			Name:        "Max-Primes",
			Value:       MAX_PRIME_SIZE,
			EnvVars:     []string{"MAX_PRIMES"},
			Destination: &flags.MaxPrimes,
		},
		&cli.BoolFlag{
			Name:        "Logger-For-Prime-Numbers",
			Value:       true,
			EnvVars:     []string{"PRIME_LOGGER"},
			Destination: &flags.PrimeLogger,
		},
		&cli.StringFlag{
			Name:        "Log-Level-For-Entire-App",
			Value:       "DEBUG",
			EnvVars:     []string{"LOG_LEVEL"},
			Destination: &flags.LogLevel,
		},
		&cli.StringFlag{
			Name:        "Server-Version",
			Value:       "1.0",
			EnvVars:     []string{"VERSION"},
			Destination: &flags.Version,
		},
	}
}
