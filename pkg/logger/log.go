package logger

type Logger interface {
	LogInfo(msg string)
	LogError(err string)
	LogWarn(warn string)
	LogDebug(msg string)
}
