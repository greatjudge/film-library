package logger

import (
	"net/http"

	"go.uber.org/zap"
)

type Logger struct {
	logger *zap.SugaredLogger
}

func NewLogger(logger *zap.SugaredLogger) *Logger {
	return &Logger{
		logger: logger,
	}
}

func (l Logger) IncomingLog(r *http.Request) {
	l.logger.Infof(
		`incoming, host: %v, url: %v, method: %v`,
		r.Host,
		r.URL,
		r.Method,
	)
}

func (l Logger) Error(mes string, where string, err error) {
	l.logger.Errorf(
		`message: "%v", where: "%v", err: "%w"`,
		mes,
		where,
		err,
	)
}

func (l Logger) LogResponse(r *http.Request, mes string, where string, status int) {
	l.logger.Infof(
		`url: %v, status-code: %d, message: "%v", where: "%v"`,
		r.URL,
		status,
		mes,
		where,
	)
}
