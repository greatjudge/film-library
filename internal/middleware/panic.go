package middleware

import (
	"filmlibr/internal/logger"
	"filmlibr/internal/sending"
	"net/http"
)

func Panic(logger logger.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				logger.Error("recovered", "panic middleware", err.(error))
				sending.JSONError(w, ErrInternal, http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
