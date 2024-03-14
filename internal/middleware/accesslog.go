package middleware

import (
	"filmlibr/internal/logger"
	"net/http"
)

func AccessLog(logger logger.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.IncomingLog(r)

		next.ServeHTTP(w, r)
	})
}
