package middleware

import (
	"errors"
	"filmlibr/internal/sending"
	"net/http"
)

var ErrBadContentType = errors.New("wrong Content-Type, must be application/json")

func JSONContentCheck(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "application/json" {
			sending.JSONError(w, ErrBadContentType, http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	})
}
