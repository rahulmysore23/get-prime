package logrus

import (
	"os"

	"github.com/rahulmysore23/get-prime/pkg/utilities"
	log "github.com/sirupsen/logrus"
)

type Logrus struct {
	Log log.Logger
}

func (l Logrus) LogDebug(msg string) {
	log.Debug(msg)
}

func (l Logrus) LogInfo(msg string) {
	log.Info(msg)
}

func (l Logrus) LogError(err string) {
	log.Error(err)
}

func (l Logrus) LogWarn(warn string) {
	log.Warn(warn)
}

func SetLogrusLevel(level string) log.Level {
	if level == utilities.PROD {
		return log.InfoLevel
	}
	return log.DebugLevel
}

func NewLogrus(logLevel string) Logrus {

	level := SetLogrusLevel(logLevel)

	var logger = log.Logger{
		Out:       os.Stderr,
		Formatter: new(log.JSONFormatter),
		Level:     level,
	}

	return Logrus{
		Log: logger,
	}
}
